package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/noisyboy-9/data_extractor/internal/config"
	"github.com/noisyboy-9/data_extractor/internal/log"
	"github.com/noisyboy-9/data_extractor/internal/service"
)

func InitApp() {
	config.LoadViper()
	config.Init()
	log.Init()
	service.Init()
}

func SetupGracefulShutdown() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.App.Info("Exit signal has been received. Shutting down . . .")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	terminateApp(ctx)
}

func terminateApp(cancelCtx context.Context) {
	service.Terminate(cancelCtx)
}
