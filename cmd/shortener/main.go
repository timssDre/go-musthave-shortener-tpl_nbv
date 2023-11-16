package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/app"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
)

func main() {
	addrConfig := config.InitConfig()
	app.Start(addrConfig)
}
