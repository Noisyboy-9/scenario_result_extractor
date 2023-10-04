package main

import (
	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/noisyboy-9/golang_api_template/internal/service"
)

func main() {
	config.LoadViper()
	config.Init()
	log.Init()
	service.Init()
}
