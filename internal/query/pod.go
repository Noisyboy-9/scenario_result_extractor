package query

import (
	"context"
	"fmt"
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
