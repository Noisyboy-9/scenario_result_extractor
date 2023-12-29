package service

import "context"

func Init() {
	InitPrometheus()
}

func Terminate(cancelCtx context.Context) {
}
