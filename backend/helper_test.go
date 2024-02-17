package main

import (
	"backend/api"
	"backend/model"
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const (
	database   = "integration"
	hmacSecret = "gtMtzZVY5n+wM/B1rUZFRz9uiUQN28C2DaKJMtdwuYqkONFID/yYiFeYYYU+l/fazLD0/DNrKw03cK4AaTcZhQ=="
)

func loadConfig() (model.Configuration, error) {
	config, err := model.LoadConfig(context.Background())
	config.DBDatabase = database
	config.HMACSecret = hmacSecret

	return config, err
}

func cleanTables(config model.Configuration) error {
	db, err := sql.Open("postgres", config.DSN())
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("TRUNCATE member_configs RESTART IDENTITY")
	return err
}

func adminToken() string {
	return getToken(model.RoleAdmin)
}

func getToken(role model.Role) string {
	jwt, _ := api.CreateJWT(
		model.LoginUser{
			Username: "user",
			Role:     role,
		},
		hmacSecret,
	)
	return jwt
}

func RemoveTimestamps(t *testing.T, body []byte) string {
	var parsed map[string]any
	err := json.Unmarshal(body, &parsed)
	require.NoError(t, err)

	createdAt, err := time.Parse(time.RFC3339, parsed["createdAt"].(string))
	require.NoError(t, err)
	updatedAt, err := time.Parse(time.RFC3339, parsed["updatedAt"].(string))
	require.NoError(t, err)

	now := time.Now()
	require.Less(t, now.Sub(createdAt), time.Minute)
	require.Less(t, now.Sub(updatedAt), time.Minute)

	delete(parsed, "createdAt")
	delete(parsed, "updatedAt")

	marshal, err := json.Marshal(parsed)
	require.NoError(t, err)
	return string(marshal)
}
