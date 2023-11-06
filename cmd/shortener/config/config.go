package config

type AddrConfig struct {
	ServerAddr string
	BaseAddr   string
}

func InitConfig() *AddrConfig {
	config := &AddrConfig{
		ServerAddr: "localhost:8080",
		BaseAddr:   "localhost:8080",
	}
	return config
}
