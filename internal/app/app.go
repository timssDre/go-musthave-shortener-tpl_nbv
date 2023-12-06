package app

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/dump"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"log"
)

type App struct {
	storageInstance *storage.Storage
	config          *config.Config
}

func NewApp(storageInstance *storage.Storage, config *config.Config) *App {
	return &App{
		storageInstance: storageInstance,
		config:          config,
	}
}

func (a *App) Start() {
	err := dump.FillFromStorage(a.storageInstance, a.config.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = api.StartRestAPI(a.config.ServerAddr, a.config.BaseURL, a.config.LogLevel, a.config.DbPath, a.storageInstance)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Stop() {
	err := dump.Set(a.storageInstance, a.config.FilePath, a.config.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

}
