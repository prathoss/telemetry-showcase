package config

import (
	"fmt"
	"os"
)

type Config struct {
	Address      string
	UsersAddress string
	BikesAddress string
	RidesAddress string
}

func NewConfig() (Config, error) {
	address := os.Getenv("SERVICE_ADDRESS")
	if address == "" {
		address = ":8080"
	}

	usersAddressKey := "USERS_ADDRESS"
	usersAddress := os.Getenv(usersAddressKey)
	if usersAddress == "" {
		return Config{}, fmt.Errorf("missing environment variable %s", usersAddressKey)
	}

	bikesAddressKey := "BIKES_ADDRESS"
	bikesAddress := os.Getenv(bikesAddressKey)
	if bikesAddress == "" {
		return Config{}, fmt.Errorf("missing environment variable %s", bikesAddressKey)
	}

	ridesAddressKey := "RIDES_ADDRESS"
	ridesAddress := os.Getenv(ridesAddressKey)
	if ridesAddress == "" {
		return Config{}, fmt.Errorf("missing environment variable %s", ridesAddressKey)
	}

	return Config{
		Address:      address,
		UsersAddress: usersAddress,
		BikesAddress: bikesAddress,
		RidesAddress: ridesAddress,
	}, nil
}
