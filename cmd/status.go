package cmd

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/noisyboy-9/data_extractor/internal/app"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/model"
	"github.com/noisyboy-9/data_extractor/internal/query"
	"github.com/noisyboy-9/data_extractor/internal/service"
	"github.com/spf13/cobra"
)

type Status struct {
	HPAs         []model.Hpa
	PodPlacement map[string][]string
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get status of the system",
	Run:   statusRunner,
}

func init() {
	rootCmd.AddCommand(statusCmd)
}

func statusRunner(cmd *cobra.Command, args []string) {
	app.InitApp()

	start := time.Now().Add(-5 * time.Minute)
	end := time.Now()
	namespace := "kube-schedule"

	log.App.Info("get hpa status ...")
	HPAs := query.GetHpaStatus(start, end, namespace)
	log.App.WithField("hpas", HPAs).Info("hpa status fetched")

	log.App.Info("get pod status ...")
	podStatus := query.GetPodStatus(start, end, namespace)
	log.App.WithField("pod_status", podStatus).Info("pod status fetched")

	timestamps := make([]time.Time, 0)
	for timestamp := range HPAs {
		timestamps = append(timestamps, timestamp)
	}
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})

	finalReport := make(map[string]Status)
	for _, timestamp := range timestamps {
		relativeTimestamp := timestamp.Sub(start).Round(time.Second)

		podNodePlacement := podStatus[timestamp]
		hpaStatus := HPAs[timestamp]

		finalReport[relativeTimestamp.String()] = Status{
			HPAs:         hpaStatus,
			PodPlacement: podNodePlacement,
		}
	}

	indentedReportJson, err := json.MarshalIndent(finalReport, "", "    ")
	if err != nil {
		log.App.WithError(err).Panic("can't marshal final report to json")
	}

	if err := service.Reporter.SaveReportToFile(indentedReportJson, start, end); err != nil {
		log.App.WithError(err).Panic("error in saving report")
	}
}
