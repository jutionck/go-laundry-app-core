package main

import (
	"github.com/jutionck/go-laundry-app-core/config"
	"log"
)

func main() {
	cfg := config.NewConfig()
	_ = cfg.DbConn()
	defer func(cfg *config.Config) {
		err := cfg.DbClose()
		if err != nil {
			log.Println(err.Error())
		}
	}(&cfg)
}
