package config

type AddrConfig struct {
	ServerAddr string
	BaseURL    string
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8080",
		BaseURL:    "http://localhost:8080",
	}
	return config
}
