package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/adapter/myapp"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
)

func main() {
	addrConfig := config.InitConfig()

	app := struct {
		myapp *myapp.Myapp
	}{}

	app.myapp = myapp.New(addrConfig.ServerAddr, addrConfig.BaseURL)

	err := app.myapp.Start()
	if err != nil {
		panic(err)
	}
}
