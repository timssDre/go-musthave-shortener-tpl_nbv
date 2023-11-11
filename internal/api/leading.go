package api

import "github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/inmemory"

type Screwdriver struct {
	*inmemory.Storage
}

func Start(ServerAddr string, BaseURL string) {
	app := struct {
		api Screwdriver
	}{}

	app.api.Storage = inmemory.New(BaseURL)

	err := app.api.StartService(ServerAddr)
	if err != nil {
		panic(err)
	}
}
