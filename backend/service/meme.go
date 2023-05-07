package service

import (
	"backend/model"
	"strings"
)

func specialize(config model.NetworkConfig) model.NetworkConfig {
	if strings.ToLower(config.Firstname) == "phillip" {
		config.Firstname += " \U0001F6BF"
	}

	return config
}
