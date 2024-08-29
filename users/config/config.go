package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Address    string
	DbConnStr  string
	RedisAddrs []string
}

func NewConfig() (Config, error) {
	address := os.Getenv("GRPC_SERVICE_ADDRESS")
	if address == "" {
		address = ":50051"
	}

	dbConnStrKey := "DB_CONNECTION_STRING"
	dbConnStr := os.Getenv(dbConnStrKey)
	if dbConnStr == "" {
		return Config{}, fmt.Errorf("could not get required environment variable: %s", dbConnStrKey)
	}

	redisAddrKey := "REDIS_ADDRESSES"
	redisAddrVal := os.Getenv(redisAddrKey)
	if redisAddrVal == "" {
		return Config{}, fmt.Errorf("could not get required environment variable: %s", redisAddrKey)
	}
	redisAddrs := strings.Split(redisAddrVal, ",")

	return Config{
		Address:    address,
		DbConnStr:  dbConnStr,
		RedisAddrs: redisAddrs,
	}, nil
}
