package query

import (
	"context"
	"fmt"
	"time"

	"github.com/noisyboy-9/data_extractor/internal/config"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/model"
	"github.com/noisyboy-9/data_extractor/internal/service"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promModel "github.com/prometheus/common/model"
)

func GetHpaStatus(start time.Time, end time.Time, namespace string) map[time.Time][]model.Hpa {
	ctx, cancel := context.WithTimeout(context.Background(), config.Prometheus.Timeout)
	defer cancel()

	r := v1.Range{
		Start: start,
		End:   end,
		Step:  config.Prometheus.Step,
	}

	query := fmt.Sprintf("kube_horizontalpodautoscaler_status_target_metric{namespace='%s'}", namespace)
	results, warnings, err := service.Promtheus.Api.QueryRange(ctx, query, r)

	if err != nil {
		log.App.WithError(err).Panic("error in getting hpa status form prometheus")
	}

	if len(warnings) > 0 {
		log.App.Warnf("hpa list warnings: %v", warnings)
	}

	result := make(map[time.Time][]model.Hpa)
	samples := results.(promModel.Matrix)
	for _, sample := range samples {
		for _, value := range sample.Values {
			// fmt.Println(value.Timestamp.Time().Format("2006-01-02 15:04:05"))
			// fmt.Println(value.Value)
			// fmt.Println(sample.Metric["horizontalpodautoscaler"])
			// fmt.Println("-------------")

			hpas := result[value.Timestamp.Time()]
			hpas = append(hpas, model.Hpa{
				MetricValue: float64(value.Value),
				Name:        string(sample.Metric["horizontalpodautoscaler"]),
			})
			result[value.Timestamp.Time()] = hpas
		}
	}

	return result
}
