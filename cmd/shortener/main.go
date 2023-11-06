package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/cmd/shortener/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/adapter/server"
)

func main() {
	addrConfig := config.InitConfig()

	flag.StringVar(&addrConfig.ServerAddr, "a", addrConfig.ServerAddr, "address and port to run server")
	flag.StringVar(&addrConfig.BaseURL, "b", addrConfig.BaseURL, "address and port to run server addrResPos")
	flag.Parse()

	err := env.Parse(addrConfig)
	if err != nil {
		panic(err)
	}

	app := struct {
		server *server.Server
	}{}

	app.server = server.New(addrConfig.ServerAddr, addrConfig.BaseURL)

	err = app.server.Start()
	if err != nil {
		panic(err)
	}
}
