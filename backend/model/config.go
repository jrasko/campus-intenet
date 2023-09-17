package model

import (
	"context"
	"fmt"

	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	Username string `env:"LOGIN_USER,required"`
	Password string `env:"LOGIN_PASSWORD_HASH,required"`
	Salt     string `env:"SALT,required"`

	ArgonTime    uint32 `env:"ARGON_TIME,required"`
	ArgonMemory  uint32 `env:"ARGON_MEMORY,required"`
	ArgonThreads uint8  `env:"ARGON_THREADS,required"`
	ArgonKeyLen  uint32 `env:"ARGON_KEY_LENGTH,required"`

	HMACSecret string `env:"HMAC_SECRET,required"`

	DBHost     string `env:"POSTGRES_HOST,default=dhcp_db"`
	DBDatabase string `env:"POSTGRES_DB,default=postgres"`
	DBUser     string `env:"POSTGRES_USER,default=postgres"`
	DBPassword string `env:"POSTGRES_PASSWORD,required"`

	URL string `env:"URL,default=:8080"`

	CIDR                 string `env:"CIDR,required"`
	OutputFile           string `env:"OUTPUT_FILE,default=user-list.json"`
	SkipDhcpNotification bool   `env:"SKIP_DHCP_RELOAD,default=false"`
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

func LoadConfig(ctx context.Context) (Configuration, error) {
	var config Configuration
	err := envconfig.Process(ctx, &config)
	if err != nil {
		return Configuration{}, err
	}

	return config, nil
}
