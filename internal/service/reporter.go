package service

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type reporter struct {
	reportFolderPath string
}

var Reporter *reporter

func InitReporter() {
	Reporter = new(reporter)
	Reporter.reportFolderPath = "reports"
}

func (reporter *reporter) SaveReportToFile(report []byte, start time.Time, end time.Time, namespace string) error {
	reportTimestamp := time.Now().Format("2006-01-02")
	if err := reporter.ensureDirectory(namespace, reportTimestamp); err != nil {
		return err
	}

	reportSavingPath := fmt.Sprintf(
		"reports/%s/%s/status_%s_%s.json",
		namespace,
		reportTimestamp,
		start.Format("2006-01-02-15:04:05"),
		end.Format("2006-01-02-15:04:05"),
	)

	return os.WriteFile(reportSavingPath, report, 0644)
}

func (reporter *reporter) ensureDirectory(namespace string, path string) error {
	reportPath := filepath.Join(reporter.reportFolderPath, namespace, path)
	return os.MkdirAll(reportPath, os.ModePerm)
}
