package app

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"log"
)

func Start(config *config.Config) {
	storageInstance := storage.NewStorage()
	err := api.StartRestAPI(config.ServerAddr, config.BaseURL, config.LogLevel, config.FilePath, storageInstance)
	if err != nil {
		log.Fatal(err)
	}
}
