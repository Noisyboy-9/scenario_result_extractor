package cmd

import (
	"encoding/json"
	"time"

	"github.com/noisyboy-9/data_extractor/internal/app"
	"github.com/noisyboy-9/data_extractor/internal/enum"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/model"
	"github.com/noisyboy-9/data_extractor/internal/query"
	"github.com/noisyboy-9/data_extractor/internal/service"
	"github.com/noisyboy-9/data_extractor/internal/util"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of the system",
	Run:   statusRunner,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

const (
	REPORT_NAMESPACE = enum.ECMUS_NAMESPACE
	REPORT_START     = "2023-12-30 01:10:30"
	REPORT_END       = "2023-12-30 01:32:30"
)

func statusRunner(cmd *cobra.Command, args []string) {
	app.InitApp()
	start, end, err := util.SetReportStartAndEndTime(REPORT_START, REPORT_END)
	if err != nil {
		log.App.WithError(err).Panic("error in getting report start and end time")
	}

	log.App.Info("get hpa status ...")
	HPAs := query.GetHpaStatus(start, end, REPORT_NAMESPACE)
	log.App.WithField("hpas", HPAs).Info("hpa status fetched")

	log.App.Info("get pod status ...")
	podStatus := query.GetPodStatus(start, end, REPORT_NAMESPACE)
	log.App.WithField("pod_status", podStatus).Info("pod status fetched")

	soretdTimestamps := util.GetSortedTimestamps(HPAs)

	finalReport := make([]model.StatusReport, len(soretdTimestamps))
	for i, timestamp := range soretdTimestamps {
		relativeTimestamp := timestamp.Sub(start).Round(time.Second)

		podNodePlacement := podStatus[timestamp]
		hpaStatus := HPAs[timestamp]

		finalReport[i] = model.StatusReport{
			Timestamp:    relativeTimestamp.String(),
			HPAs:         hpaStatus,
			PodPlacement: podNodePlacement,
		}
	}

	indentedReportJson, err := json.MarshalIndent(finalReport, "", "    ")
	if err != nil {
		log.App.WithError(err).Panic("can't marshal final report to json")
	}

	if err := service.Reporter.SaveReportToFile(indentedReportJson, start, end, REPORT_NAMESPACE); err != nil {
		log.App.WithError(err).Panic("error in saving report")
	}
}
