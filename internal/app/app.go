package app

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/dump"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"log"
)

func Start(config *config.Config, storageInstance *storage.Storage) {
	err := dump.FillFromStorage(storageInstance, config.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = api.StartRestAPI(config.ServerAddr, config.BaseURL, config.LogLevel, storageInstance)
	if err != nil {
		log.Fatal(err)
	}
}

func Stop(config *config.Config, storageInstance *storage.Storage) {
	err := dump.Set(storageInstance, config.FilePath, config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

}
