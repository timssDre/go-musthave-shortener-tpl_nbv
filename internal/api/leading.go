package api

import "github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/server"

type Screwdriver struct {
	*server.Storage
}

func Start(ServerAddr string, BaseURL string) {
	app := struct {
		api Screwdriver
	}{}

	app.api.Storage = server.New(BaseURL)

	err := app.api.StartService(ServerAddr)
	if err != nil {
		panic(err)
	}
}
