package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
)

func main() {
	addrConfig := config.InitConfig()

	app := struct {
		api *api.Storage
	}{}

	app.api = api.New(addrConfig.BaseURL)

	err := app.api.Start(addrConfig.ServerAddr)
	if err != nil {
		panic(err)
	}
}
