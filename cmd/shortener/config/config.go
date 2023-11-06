package config

type AddrConfig struct {
	ServerAddr string
	BaseUrl    string
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8080",
		BaseUrl:    "default",
	}
	return config
}
