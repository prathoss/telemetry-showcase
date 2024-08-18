package config

import "os"

type Config struct {
	Address string
}

func NewConfig() Config {
	address := os.Getenv("SERVICE_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	return Config{
		Address: address,
	}
}
