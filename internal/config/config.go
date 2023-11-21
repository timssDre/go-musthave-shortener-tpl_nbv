package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
	LogLevel   string `env:"FLAG_LOG_LEVEL"`
}

func InitConfig() *Config {
	config := &Config{
		ServerAddr: "localhost:8080",
		BaseURL:    "http://localhost:8080",
		LogLevel:   "info",
	}
	flag.StringVar(&config.ServerAddr, "a", config.ServerAddr, "address and port to run api")
	flag.StringVar(&config.BaseURL, "b", config.BaseURL, "address and port to run api addrResPos")
	flag.StringVar(&config.LogLevel, "c", config.LogLevel, "log level")
	flag.Parse()
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
