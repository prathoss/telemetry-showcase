package config

import "os"

type Config struct {
	Address string
}

func NewConfig() Config {
	address := os.Getenv("GRPC_SERVICE_ADDRESS")
	if address == "" {
		address = ":50051"
	}

	return Config{
		Address: address,
	}
}
