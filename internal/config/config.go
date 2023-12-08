package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddr  string `env:"SERVER_ADDRESS"`
	BaseURL     string `env:"BASE_URL"`
	LogLevel    string `env:"FLAG_LOG_LEVEL"`
	FilePath    string `env:"FILE_STORAGE_PATH"`
	DBPath      string `env:"db"`
	DataBaseDSN string `env:"DATABASE_DSN"`
}

func InitConfig() *Config {
	config := &Config{
		ServerAddr:  "localhost:8080",
		BaseURL:     "http://localhost:8080",
		LogLevel:    "info",
		FilePath:    "short-url-db.json",
		DBPath:      fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "nbvpass", "postgres"),
		DataBaseDSN: "",
	}

	flag.StringVar(&config.ServerAddr, "a", config.ServerAddr, "address and port to run api")
	flag.StringVar(&config.BaseURL, "b", config.BaseURL, "address and port to run api addrResPos")
	flag.StringVar(&config.LogLevel, "c", config.LogLevel, "log level")
	flag.StringVar(&config.DataBaseDSN, "e", config.DataBaseDSN, "address to base store in-memory")
	flag.StringVar(&config.DBPath, "d", config.DBPath, "address to base store in-memory")
	//flag.StringVar(&config.DataBaseDSN, "d", config.DataBaseDSN, "address to base store in-memory")
	//flag.StringVar(&config.DBPath, "e", config.DBPath, "address to base store in-memory")
	flag.StringVar(&config.FilePath, "f", config.FilePath, "address to file in-memory")

	flag.Parse()
	err := env.Parse(config)
	if err != nil {
		panic(err)
	}

	return config
}
