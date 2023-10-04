package log

import (
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	HttpServerLogPrefix = "httpServer"
)

var (
	App *logrus.Logger
)

func Init() {
	App := logrus.New()

	App.SetFormatter(&runtime.Formatter{
		ChildFormatter: &logrus.JSONFormatter{},
		Line:           true,
		File:           true,
		Package:        false,
		BaseNameOnly:   false,
	})

	configPrefix := "logging.app"

	if viper.GetBool(configPrefix + "stdout") {
		App.SetOutput(os.Stdout)
	}

	level := logrus.DebugLevel
	if viper.GetString("app.env") == "production" {
		var err error
		level, err = logrus.ParseLevel(viper.GetString("logging.app.level"))
		if err != nil {
			panic(err)
		}
	}
	App.SetLevel(level)
}
