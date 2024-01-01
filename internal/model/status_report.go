package model

type StatusReport struct {
	Timestamp    string              `json:"timestamp" default:"timestamp"`
	HPAs         []Hpa               `json:"hpa" default:"hpa"`
	PodPlacement map[string][]string `json:"pod_placement" default:"pod_placement"`
}
