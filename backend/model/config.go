package model

import (
	"context"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	Username     string `env:"LOGIN_USER,required"`
	PasswordHash string `env:"LOGIN_PASSWORD_HASH,required"`
	HMACSecret   string `env:"HMAC_SECRET,required"`

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
	_, _, _, err = argon2id.DecodeHash(config.PasswordHash)
	if err != nil {
		return Configuration{}, fmt.Errorf("when reading password hash: %w", err)
	}

	return config, nil
}
