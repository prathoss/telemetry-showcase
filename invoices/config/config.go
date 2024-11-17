package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	KafkaBrokers []string
	RidesAddress string
}

func NewConfig() (Config, error) {
	envBrokersKey := "KAFKA_BROKERS"
	envBrokers := os.Getenv("KAFKA_BROKERS")
	if envBrokers == "" {
		return Config{}, fmt.Errorf("could not find %s environment variable", envBrokersKey)
	}
	brokers := strings.Split(envBrokers, ",")

	ridesAddressKey := "RIDES_ADDRESS"
	ridesAddress := os.Getenv(ridesAddressKey)
	if ridesAddress == "" {
		return Config{}, fmt.Errorf("could not find %s environment variable", ridesAddressKey)
	}

	return Config{
		KafkaBrokers: brokers,
		RidesAddress: ridesAddress,
	}, nil
}
