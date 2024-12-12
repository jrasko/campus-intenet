package repository

import (
	"backend/model"
	"context"

	"gorm.io/gorm"
)

func (t *RepositoryTest) TestRepository_GetNetConfig() {
	config, err := t.repository.GetNetConfig(context.Background(), member1.NetConfigID)
	t.Require().NoError(err)
	t.NetEqual(member1.NetConfig, config)
}

func (t *RepositoryTest) TestRepository_ListNetConfig() {
	config, err := t.repository.ListNetConfigs(context.Background(), model.NetworkRequestParams{})
	t.Require().NoError(err)
	t.Len(config, 3)
	t.NetEqual(config[0], member2.NetConfig)
	t.NetEqual(config[1], net1)
	t.NetEqual(config[2], member1.NetConfig)
}

func (t *RepositoryTest) TestRepository_CreateNetConfig() {
	cfg := model.NetConfig{
		Name:         "some name",
		Mac:          "AA:AA:AA:BB:BB:BB",
		IP:           "11.22.33.44",
		Manufacturer: "manu",
		Disabled:     true,
	}
	fromDB, err := t.repository.CreateOrUpdateNetConfig(context.Background(), cfg)
	t.Require().NoError(err)
	cfg.ID = fromDB.ID
	t.NetEqual(cfg, fromDB)
}

func (t *RepositoryTest) TestRepository_DeleteNetConfig() {
	err := t.repository.DeleteNetConfig(context.Background(), net1.ID)
	t.Require().NoError(err)

	_, err = t.repository.GetNetConfig(context.Background(), net1.ID)
	t.Equal(gorm.ErrRecordNotFound, err)
}

func (t *RepositoryTest) TestRepository_Deactivate() {
	err := t.repository.Deactivate(context.Background(), []int{member1.NetConfigID, member2.NetConfigID, net1.ID})
	t.Require().NoError(err)

	cfgs, err := t.repository.ListNetConfigs(context.Background(), model.NetworkRequestParams{})
	t.Require().NoError(err)
	for _, cfg := range cfgs {
		t.True(cfg.Disabled)
	}
}

func (t *RepositoryTest) TestRepository_GetAllIPs() {
	ips, err := t.repository.GetAllIPs(context.Background())
	t.Require().NoError(err)
	t.Len(ips, 3)
	t.Contains(ips, member1.NetConfig.IP)
	t.Contains(ips, member2.NetConfig.IP)
	t.Contains(ips, net1.IP)
}
