package repository

import (
	"backend/model"
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func setupDB() (NetworkRepository, error) {
	cfg := model.Configuration{
		DBHost:     "localhost",
		DBDatabase: "network",
		DBUser:     "network",
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
	repo, err := New(cfg.DSN())
	if err != nil {
		return NetworkRepository{}, err
	}
	repo.db.Exec("TRUNCATE member_configs RESTART IDENTITY")
	return repo, err
}

func TestNew(t *testing.T) {
	repo, err := New("")
	assert.Error(t, err)
	assert.Equal(t, NetworkRepository{}, repo)
}

func TestNetworkRepository(t *testing.T) {
	repo, creationErr := setupDB()
	require.NoError(t, creationErr)

	config := model.MemberConfig{
		Firstname: "first",
		Lastname:  "name",
		Mac:       "00:11:22:33:44:55",
		RoomNr:    "00-11",
		HasPaid:   false,
		WG:        "25a",
		Email:     "bernd-das-brot@alumni.test-provider.com",
		Phone:     "012345678901",
		IP:        "192.168.1.1",
	}
	config2 := model.MemberConfig{
		Firstname: "mister",
		Lastname:  "x",
		Mac:       "aa:aa:aa:aa:aa:aa",
		RoomNr:    "99-99",
		HasPaid:   false,
		WG:        "13x",
		Email:     "mail@email.email",
		Phone:     "0999888777666",
		IP:        "10.0.0.1",
	}
	updatedConfig := model.MemberConfig{
		Firstname: "Bernd",
		Lastname:  "Das Brot",
		Mac:       "55:44:33:22:11:00",
		RoomNr:    "11-00",
		HasPaid:   true,
		WG:        "29b",
		Email:     "other-email@alumni.test-provider.com",
		Phone:     "987654321",
		IP:        "172.0.0.1",
	}

	t.Run("it creates a member", func(t *testing.T) {
		newConfig, err := repo.UpdateNetworkConfig(ctx, config)
		assert.NoError(t, err)
		assert.NotEmpty(t, newConfig.ID)
		config.ID = newConfig.ID
		assert.Equal(t, config, newConfig)
	})
	t.Run("it creates another member", func(t *testing.T) {
		newConfig, err := repo.UpdateNetworkConfig(ctx, config2)
		config2.ID = newConfig.ID
		assert.NoError(t, err)
		assert.Equal(t, config2, newConfig)
	})
	t.Run("it retrevies a single member", func(t *testing.T) {
		cfg, err := repo.GetNetworkConfig(ctx, config.ID)
		assert.NoError(t, err)
		assert.Equal(t, cfg, config)
	})
	t.Run("it retreives multiple members", func(t *testing.T) {
		cfgs, err := repo.GetAllNetworkConfigs(ctx)
		assert.NoError(t, err)
		assert.Contains(t, cfgs, config)
		assert.Contains(t, cfgs, config2)
	})
	t.Run("it retreives all ips", func(t *testing.T) {
		ips, err := repo.GetAllIPs(ctx)
		assert.NoError(t, err)
		assert.Contains(t, ips, config.IP)
		assert.Contains(t, ips, config2.IP)
	})
	t.Run("it updates a member", func(t *testing.T) {
		updatedConfig.ID = config.ID
		newConfig, err := repo.UpdateNetworkConfig(ctx, updatedConfig)
		assert.NoError(t, err)
		assert.Equal(t, updatedConfig, newConfig)
	})
	t.Run("it resets payments", func(t *testing.T) {
		err := repo.ResetPayment(ctx)
		assert.NoError(t, err)

		cfgs, err := repo.GetAllNetworkConfigs(ctx)
		assert.NoError(t, err)
		for _, cfg := range cfgs {
			assert.False(t, cfg.HasPaid)
		}
	})
	t.Run("it deletes a member", func(t *testing.T) {
		err := repo.DeleteNetworkConfig(ctx, config2.ID)
		assert.NoError(t, err)

		_, err = repo.GetNetworkConfig(ctx, config2.ID)
		assert.ErrorIs(t, err, ErrNotFound)
	})

}
