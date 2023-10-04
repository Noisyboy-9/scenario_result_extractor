package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/log"
	"github.com/noisyboy-9/golang_api_template/internal/service"
)

func main() {
	config.LoadViper()
	config.Init()
	log.Init()
	service.Init()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.App.Info("Exit signal has been received. Shutting down . . .")
}
