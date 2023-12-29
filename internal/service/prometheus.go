package service

import (
	"github.com/noisyboy-9/data_extractor/internal/config"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	promconfig "github.com/prometheus/common/config"
)

type prometheus struct {
	Client api.Client
	Api    v1.API
}

var Promtheus *prometheus

func InitPrometheus() {
	log.App.Info("init connection to prometheus . . .")

	Promtheus = new(prometheus)

	client, err := api.NewClient(api.Config{
		Address:      config.Prometheus.Address,
		RoundTripper: promconfig.NewBasicAuthRoundTripper(config.Prometheus.Username, promconfig.Secret(config.Prometheus.Password), "", "", api.DefaultRoundTripper),
	})

	if err != nil {
		log.App.Panicln(err.Error())
	}

	Promtheus.Client = client
	Promtheus.Api = v1.NewAPI(client)
}
