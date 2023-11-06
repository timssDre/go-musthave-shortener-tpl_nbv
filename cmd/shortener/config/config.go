package config

type AddrConfig struct {
	ServerAddr string `env:"SERVER_ADDRESS"`
	BaseURL    string `env:"BASE_URL"`
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8080",
		BaseURL:    "http://localhost:8080",
	}
	return config
}
