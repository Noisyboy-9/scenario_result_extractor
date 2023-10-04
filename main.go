package main

import "github.com/noisyboy-9/golang_api_template/internal/config"

func main() {
	config.LoadViper()
	config.Init()
}
