package cmd

import (
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

	log.App.Info("get node list ... ")
	nodes := query.GetNodeList()
	log.App.WithField("nodes", nodes).Info("node list fetched")
}
