package util

import (
	"sort"
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

	end, err = time.ParseInLocation("2006-01-02 15:04:05", endStr, tehranTimezone)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}

	return start, end, nil
}

func GetSortedTimestamps[T any](datas map[time.Time]T) []time.Time {
	timestamps := make([]time.Time, 0)
	for timestamp := range datas {
		timestamps = append(timestamps, timestamp)
	}
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})
	return timestamps
}
