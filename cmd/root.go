package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Extract paper's required data from prometheus and do something with it",
}

var (
	ReportNamespace    string
	ReportStart        string
	ReportEnd          string
	ReportScenarioName string
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&ReportStart, "start", "", "scenario start time with format 2006-01-02 15:04:05")
	rootCmd.PersistentFlags().StringVar(&ReportEnd, "end", "", "scenario end time with format 2006-01-02 15:04:05")
	rootCmd.PersistentFlags().StringVar(&ReportNamespace, "namespace", "", "scenario namespace")
	rootCmd.PersistentFlags().StringVar(&ReportScenarioName, "name", "", "scenario name")
}
