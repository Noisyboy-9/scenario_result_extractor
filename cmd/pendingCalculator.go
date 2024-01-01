package cmd

import (
	"encoding/json"

	"github.com/noisyboy-9/data_extractor/internal/app"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/query"
	"github.com/noisyboy-9/data_extractor/internal/service"
	"github.com/noisyboy-9/data_extractor/internal/util"
	"github.com/spf13/cobra"
)

// pendingCalculatorCmd represents the pendingCalculator command
var pendingCalculatorCmd = &cobra.Command{
	Use:   "pendingCalculator",
	Short: "calculates the amount of time each pod is in pending state",
	Run:   pendingDurationCalculatorCmd,
}

func init() {
	rootCmd.AddCommand(pendingCalculatorCmd)
}

func pendingDurationCalculatorCmd(*cobra.Command, []string) {
	app.InitApp()

	start, end, err := util.SetReportStartAndEndTime(ReportStart, ReportEnd)
	if err != nil {
		log.App.WithError(err).Panic("error in getting report start and end time")
	}

	allPods := query.GetPodList(start, end, ReportNamespace)
	podToStartTimeMap, readyPods := query.GetPodReadyDuration(start, end, ReportNamespace)
	createdButNotReadyPods := util.MakeUnique[[]string](util.GetSetDiff[[]string](allPods, readyPods))

	resultJSONMap := map[string]interface{}{
		"ready_pod_duration":         podToStartTimeMap,
		"created_but_not_ready_pods": createdButNotReadyPods,
	}

	indentedReportJson, err := json.MarshalIndent(resultJSONMap, "", "    ")
	if err != nil {
		log.App.WithError(err).Panic("can't marshal final report to json")
	}

	if err := service.Reporter.SaveReportToFile(indentedReportJson, ReportScenarioName, ReportNamespace, "pending"); err != nil {
		log.App.WithError(err).Panic("error in saving report")
	}
}
