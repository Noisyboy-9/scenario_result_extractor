package config

import "time"

type prometheus struct {
	Address  string
	Username string
	Password string
	Timeout  time.Duration
	Step     time.Duration
}

var Prometheus *prometheus
