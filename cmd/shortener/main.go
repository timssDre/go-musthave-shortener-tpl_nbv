package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/app"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

func main() {
	addrConfig := config.InitConfig()

	storageInstance := storage.NewStorage()
	appInstance := app.NewApp(storageInstance, addrConfig)
	appInstance.Start()
	appInstance.Stop()
}
