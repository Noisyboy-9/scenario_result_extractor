package cmd

import (
	"time"

	"github.com/noisyboy-9/data_extractor/internal/app"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/query"
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

func statusRunner(cmd *cobra.Command, args []string) {
	app.InitApp()

	start := time.Now().Add(-5 * time.Minute)
	end := time.Now()
	namespace := "kube-schedule"

	log.App.Info("get hpa status ...")
	hpas := query.GetHpaStatus(start, end, namespace)
	log.App.WithField("hpas", hpas).Info("hpa status fetched")

	log.App.Info("get pod status ...")
	podStatus := query.GetPodStatus(start, end, namespace)
	log.App.WithField("pod_status", podStatus).Info("pod status fetched")
}
