package main

import (
	"github.com/jutionck/go-laundry-app-core/config"
	"github.com/jutionck/go-laundry-app-core/manager"
)

func main() {
	appConfig := config.NewConfig()
	_ = manager.NewInfra(appConfig)
}
