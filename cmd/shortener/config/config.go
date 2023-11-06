package config

type AddrConfig struct {
	ServerAddr string
	BaseURL    string
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8081",
		BaseURL:    "http://localhost:8081",
	}
	return config
}
