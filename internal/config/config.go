package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type AddrConfig struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8080",
		BaseURL:    "http://localhost:8080",
	}
	flag.StringVar(&config.ServerAddr, "a", config.ServerAddr, "address and port to run api")
	flag.StringVar(&config.BaseURL, "b", config.BaseURL, "address and port to run api addrResPos")
	flag.Parse()
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
