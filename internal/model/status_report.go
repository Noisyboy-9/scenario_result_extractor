package model

type StatusReport struct {
	Timestamp    string              `json:"timestamp,omitempty" default:"timestamp"`
	HPAs         []Hpa               `json:"hpa,omitempty" default:"hpa"`
	PodPlacement map[string][]string `json:"pod_placement,omitempty" default:"pod_placement"`
}
