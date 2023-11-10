package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/adapter/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
)

func main() {
	addrConfig := config.InitConfig()

	app := struct {
		api *api.Api
	}{}

	app.api = api.New(addrConfig.ServerAddr, addrConfig.BaseURL)

	err := app.api.Start()
	if err != nil {
		panic(err)
	}
}
