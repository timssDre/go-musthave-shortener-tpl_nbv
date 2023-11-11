package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

func main() {
	addrConfig := config.InitConfig()

	app := struct {
		api *storage.Storage
	}{}

	app.api = storage.New(addrConfig.BaseURL)

	err := app.api.Start(addrConfig.ServerAddr)
	if err != nil {
		panic(err)
	}
}
