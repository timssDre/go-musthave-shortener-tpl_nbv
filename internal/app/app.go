package app

import (
	"log"

	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/dump"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/repository"
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
	db, err := repository.InitDatabase(a.config.DBPath)
	if err != nil {
		log.Fatal(err)
	}
	dbDNSTurn := true
	if a.UseDatabase() {
		err = dump.FillFromStorage(a.storageInstance, a.config.FilePath)
		if err != nil {
			log.Fatal(err)
		}
		dbDNSTurn = false
	}

	err = api.StartRestAPI(a.config.ServerAddr, a.config.BaseURL, a.config.LogLevel, db, dbDNSTurn, a.storageInstance)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) UseDatabase() bool {
	return a.config.DBPath == ""
}

func (a *App) Stop() {
	if a.UseDatabase() {
		err := dump.Set(a.storageInstance, a.config.FilePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
