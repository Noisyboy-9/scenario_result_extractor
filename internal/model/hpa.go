package model

type Hpa struct {
	MetricValue float64 `json:"metric_value" default:"metric_value"`
	Name        string  `json:"name" default:"name"`
}
