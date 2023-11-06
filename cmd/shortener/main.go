package main

import (
	"flag"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/cmd/shortener/config"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/adapter/server"
)

var (
	urlMap  = make(map[string]string)
	addr    string
	BaseURL string
)

func main() {
	addrConfig := config.InitConfig()

	flag.StringVar(&addr, "a", addrConfig.ServerAddr, "address and port to run server")
	flag.StringVar(&BaseURL, "b", addrConfig.BaseURL, "address and port to run server addrResPos")
	flag.Parse()

	app := struct {
		server *server.Server
	}{}

	app.server = server.New(addr, addrConfig.BaseURL)

	err := app.server.Start()
	if err != nil {
		panic(err)
	}
}
