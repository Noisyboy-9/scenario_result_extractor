package config

type prometheus struct {
	Address  string
	Username string
	Password string
}

var Prometheus *prometheus
