package query

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"time"

	"github.com/noisyboy-9/data_extractor/internal/config"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/service"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
)

func GetPodStatus(start time.Time, end time.Time, namespace string) map[time.Time]map[string][]string {
	ctx, cancel := context.WithTimeout(context.Background(), config.Prometheus.Timeout)
	defer cancel()

	query := fmt.Sprintf("kube_pod_info{namespace='%s'}", namespace)
	r := v1.Range{
		Start: start,
		End:   end,
		Step:  config.Prometheus.Step,
	}
	results, warnings, err := service.Promtheus.Api.QueryRange(ctx, query, r)

	if err != nil {
		log.App.WithError(err).Panic("error in getting pod status form prometheus")
	}

	if len(warnings) > 0 {
		log.App.Warnf("get pod status warnings: %v", warnings)
	}

	result := make(map[time.Time]map[string][]string)
	samples := results.(promModel.Matrix)

	for _, sample := range samples {
		for _, value := range sample.Values {
			podPlacementStatus := result[value.Timestamp.Time()]
			if podPlacementStatus == nil {
				podPlacementStatus = make(map[string][]string)
			}

			pods := podPlacementStatus[string(sample.Metric["node"])]
			pods = append(pods, string(sample.Metric["pod"]))
			podPlacementStatus[string(sample.Metric["node"])] = pods

			result[value.Timestamp.Time()] = podPlacementStatus
		}
	}

	return result
}

func GetPodList(start time.Time, end time.Time, namespace string) []string {
	ctx, cancel := context.WithTimeout(context.Background(), config.Prometheus.Timeout)
	defer cancel()

	query := fmt.Sprintf("kube_pod_created{namespace='%s'}", namespace)
	r := v1.Range{
		Start: start,
		End:   end,
		Step:  config.Prometheus.Step,
	}
	results, warnings, err := service.Promtheus.Api.QueryRange(ctx, query, r)
	if err != nil {
		log.App.WithError(err).Panic("error in getting pod status form prometheus")
	}

	if len(warnings) > 0 {
		log.App.Warnf("get pod status warnings: %v", warnings)
	}

	pods := make([]string, 0)
	samples, ok := results.(promModel.Matrix)
	if !ok {
		log.App.Info("not ok")
	}

	for _, sample := range samples {
		for labelName, labelValue := range sample.Metric {
			if labelName == "pod" {
				pods = append(pods, string(labelValue))
			}
		}
	}

	return pods
}

func GetPodReadyDuration(start time.Time, end time.Time, namespace string) (map[string]string, []string) {
	ctx, cancel := context.WithTimeout(context.Background(), config.Prometheus.Timeout)
	defer cancel()

	query := fmt.Sprintf("kube_pod_status_ready_time{namespace='%s'}", namespace)
	r := v1.Range{
		Start: start,
		End:   end,
		Step:  config.Prometheus.Step,
	}
	results, warnings, err := service.Promtheus.Api.QueryRange(ctx, query, r)

	if err != nil {
		log.App.WithError(err).Panic("error in getting pod status form prometheus")
	}

	if len(warnings) > 0 {
		log.App.Warnf("get pod status warnings: %v", warnings)
	}

	podToReadyDurationTimeMap := make(map[string]string, 0)
	readyPods := make([]string, 0)
	samples := results.(promModel.Matrix)

	for _, sample := range samples {
		for _, value := range sample.Values {
			readyDuration := time.Duration(float64(value.Value) * float64(time.Nanosecond))
			name := string(sample.Metric["pod"])

			podToReadyDurationTimeMap[name] = strconv.FormatFloat(readyDuration.Seconds(), 'f', -1, 64) + "s"

			if !slices.Contains(readyPods, name) {
				readyPods = append(readyPods, name)
			}
		}
	}

	return podToReadyDurationTimeMap, readyPods
}
