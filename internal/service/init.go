package service

import "context"

func Init() {
	InitPrometheus()
	InitReporter()
}

func Terminate(cancelCtx context.Context) {
}
