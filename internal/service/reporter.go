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

func (reporter *reporter) SaveReportToFile(report []byte, scenarioName string, namespace string, reportType string) error {
	reportTimestamp := time.Now().Format("2006-01-02")
	if err := reporter.ensureDirectory(namespace, reportTimestamp, reportType); err != nil {
		return err
	}

	reportSavingPath := fmt.Sprintf(
		"reports/%s/%s/%s/%s.json",
		namespace,
		reportTimestamp,
		reportType,
		scenarioName,
	)

	return os.WriteFile(reportSavingPath, report, 0644)
}

func (reporter *reporter) ensureDirectory(namespace string, path string, reportType string) error {
	reportPath := filepath.Join(reporter.reportFolderPath, namespace, path, reportType)
	return os.MkdirAll(reportPath, os.ModePerm)
}
