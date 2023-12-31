package util

import (
	"time"

	"github.com/noisyboy-9/data_extractor/internal/log"
)

func SetReportStartAndEndTime(startStr string, endStr string) (start time.Time, end time.Time, err error) {
	tehranTimezone, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		log.App.WithError(err).Panic("error in getting loading asia/tehran timezone")
	}

	start, err = time.ParseInLocation("2006-01-02 15:04:05", startStr, tehranTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	end, err = time.ParseInLocation("2006-01-02 15:04:05", "2023-12-30 00:43:15", tehranTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return start, end, nil
}
