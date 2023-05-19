package model

import "fmt"

type Configuration struct {
	Username   string
	Password   string
	HMACSecret string
	Salt       string
	DBHost     string
	DBDatabase string
	DBUser     string
	DBPassword string
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
