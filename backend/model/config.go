package model

import (
	"fmt"
	"os"
)

const defaultUrl = ":8080"
const defaultFile = "dhcpd.conf"

type Configuration struct {
	Username string
	Password string
	Salt     string

	HMACSecret string

	DBHost     string
	DBDatabase string
	DBUser     string
	DBPassword string

	URL string

	CIDR       string
	OutputFile string
}

func (c Configuration) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		c.DBHost,
		c.DBUser,
		c.DBPassword,
		c.DBDatabase,
	)
}

func LoadConfig() Configuration {
	config := Configuration{
		Username:   os.Getenv("LOGIN_USER"),
		Password:   os.Getenv("LOGIN_PASSWORD_HASH"),
		HMACSecret: os.Getenv("HMAC_SECRET"),
		Salt:       os.Getenv("SALT"),
		DBDatabase: os.Getenv("POSTGRES_DB"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		CIDR:       os.Getenv("CIDR"),

		OutputFile: defaultFile,
		URL:        defaultUrl,
	}
	if (config == Configuration{}) {
		panic("empty config")
	}
	return config
}
