package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/noisyboy-9/golang_api_template/internal/config"
	"github.com/noisyboy-9/golang_api_template/internal/handler"
	"github.com/noisyboy-9/golang_api_template/internal/log"
)

type httpServer struct {
	e echo.Echo
}

var HttpServer *httpServer

func InitHttpServer() {
	httpServer := new(httpServer)
	httpServer.e = *echo.New()

	httpServer.registerRoutes()

	serverUrl := fmt.Sprintf("%s:%d", config.HttpServer.Host, config.HttpServer.Port)
	if err := httpServer.e.Start(serverUrl); err != nil {
		log.App.WithField("err", err.Error()).Fatalf("can't start web server")
	}

}

func (server *httpServer) registerRoutes() {
	HttpServer.e.GET("/", handler.SayHello)
}
