package repository

import (
	"backend/model"
	"context"
)

func (t *RepositoryTest) TestRepository_GetNetConfig() {
	config, err := t.repository.GetNetConfig(context.Background(), member1.NetConfigID)
	t.Require().NoError(err)
	t.NetEqual(member1.NetConfig, config)
}

func (t *RepositoryTest) TestRepository_ListNetConfig() {
	config, err := t.repository.ListNetConfigs(context.Background(), model.NetworkRequestParams{})
	t.Require().NoError(err)
	t.Len(config, 2)
	t.NetEqual(config[0], member2.NetConfig)
	t.NetEqual(config[1], member1.NetConfig)
}
