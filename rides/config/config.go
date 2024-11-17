package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Address      string
	DbConnStr    string
	KafkaBrokers []string
	BikesAddress string
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

	kafkaBrokersKey := "KAFKA_BROKERS"
	kafkaBrokersStr := os.Getenv(kafkaBrokersKey)
	if kafkaBrokersStr == "" {
		return Config{}, fmt.Errorf("could not get required environment variable: %s", kafkaBrokersKey)
	}
	kafkaBrokers := strings.Split(kafkaBrokersStr, ",")

	bikesAddressKey := "BIKES_ADDRESS"
	bikesAddress := os.Getenv(bikesAddressKey)
	if bikesAddress == "" {
		return Config{}, fmt.Errorf("could not get required environment variable: %s", bikesAddressKey)
	}

	return Config{
		Address:      address,
		DbConnStr:    dbConnStr,
		KafkaBrokers: kafkaBrokers,
		BikesAddress: bikesAddress,
	}, nil
}
