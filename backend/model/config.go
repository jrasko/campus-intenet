package model

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	defaultUrl  = ":8080"
	defaultFile = "user-list.json"

	argonKeyLength = 64              // 512 bits
	argonThreads   = 8               // recommended: 2 x server cores
	argonMemory    = 2 * 1024 * 1024 // [in KB] - 2 GiB
	argonTime      = 4
)

type Configuration struct {
	Username string
	Password string

	Salt         string
	ArgonTime    uint32
	ArgonMemory  uint32
	ArgonThreads uint8
	ArgonKeyLen  uint32

	HMACSecret string

	DBHost     string
	DBDatabase string
	DBUser     string
	DBPassword string

	URL string

	CIDR                 string
	OutputFile           string
	SkipDhcpNotification bool
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

func LoadConfig() (Configuration, error) {
	var err error
	skipDhcpNotification := false

	if os.Getenv("SKIP_DHCP_RELOAD") != "" {
		skipDhcpNotification, err = strconv.ParseBool(os.Getenv("SKIP_DHCP_RELOAD"))
		if err != nil {
			return Configuration{}, fmt.Errorf("parsing env 'SKIP_DHCP_RELOAD': %w", err)
		}
	}

	config := Configuration{
		Username:   os.Getenv("LOGIN_USER"),
		Password:   os.Getenv("LOGIN_PASSWORD_HASH"),
		Salt:       os.Getenv("SALT"),
		HMACSecret: os.Getenv("HMAC_SECRET"),
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBDatabase: os.Getenv("POSTGRES_DB"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		CIDR:       os.Getenv("CIDR"),

		URL:        defaultUrl,
		OutputFile: defaultFile,

		ArgonTime:    argonTime,
		ArgonMemory:  argonMemory,
		ArgonThreads: argonThreads,
		ArgonKeyLen:  argonKeyLength,

		SkipDhcpNotification: skipDhcpNotification,
	}
	if (config == Configuration{}) {
		return Configuration{}, errors.New("empty config")
	}
	return config, nil
}
