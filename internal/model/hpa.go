package model

type Hpa struct {
	MetricValue float64 `json:"metric_value,omitempty" default:"metric_value"`
	Name        string  `json:"name,omitempty" default:"name"`
}
