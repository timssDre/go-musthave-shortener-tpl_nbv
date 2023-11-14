package app

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/storage"
)

func Start(config *config.AddrConfig) {
	storageInstance := storage.New()

	err := api.StartService(config.ServerAddr, config.BaseURL, storageInstance)
	if err != nil {
		panic(err)
	}
}
