package main

import (
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/api"
	"github.com/timssDre/go-musthave-shortener-tpl_nbv.git/internal/config"
)

func main() {
	addrConfig := config.InitConfig()
	api.Start(addrConfig.ServerAddr, addrConfig.BaseURL)
}
