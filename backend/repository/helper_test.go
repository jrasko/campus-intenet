package repository

import (
	"backend/model"
	"context"

	"github.com/stretchr/testify/suite"
)

var (
	member1 = model.Member{
		Firstname: "first",
		Lastname:  "name",
		NetConfig: model.NetConfig{
			Mac:          "00:11:22:33:44:55",
			IP:           "192.168.1.1",
			Manufacturer: "test",
		},
		RoomNr:  room1.Number,
		Room:    room1,
		HasPaid: false,
		Email:   "bernd-das-brot@alumni.test-provider.com",
		Phone:   "012345678901",
	}
	member2 = model.Member{
		Firstname: "mister",
		Lastname:  "x",
		NetConfig: model.NetConfig{
			Mac: "aa:aa:aa:aa:aa:aa",
			IP:  "10.0.0.1",
		},
		HasPaid: false,
		Room:    room2,
		RoomNr:  room2.Number,
		Email:   "mail@email.email",
		Phone:   "0999888777666",
	}
)

var (
	room1 = model.Room{
		Number: "1",
		WG:     "1",
		Block:  "1",
	}
	room2 = model.Room{
		Number: "2",
		WG:     "2",
		Block:  "2",
	}
	room3 = model.Room{
		Number: "3",
		WG:     "3",
		Block:  "3",
	}
	room4 = model.Room{
		Number: "4",
		WG:     "4",
		Block:  "4",
	}
)

type RepositoryTest struct {
	suite.Suite
	repository Repository
}

func (t *RepositoryTest) SetupTest() {
	var err error
	t.repository, err = setupDB()
	t.Require().NoError(err)
	t.insertSampleData()
}

func setupDB() (Repository, error) {
	cfg, err := model.LoadConfig(context.Background())
	if err != nil {
		return Repository{}, err
	}
	cfg.DBDatabase = "testing"
	repo, err := New(cfg.DSN())
	if err != nil {
		return Repository{}, err
	}
	repo.db.Exec("TRUNCATE members, rooms, net_configs RESTART IDENTITY")
	return repo, nil
}

func (t *RepositoryTest) MemberEqual(expected model.Member, actual model.Member, msgAndArgs ...interface{}) bool {
	expected.CreatedAt = actual.CreatedAt
	expected.UpdatedAt = actual.UpdatedAt
	expected.NetConfig.CreatedAt = actual.NetConfig.CreatedAt
	expected.NetConfig.UpdatedAt = actual.NetConfig.UpdatedAt
	expected.NetConfig.ID = actual.NetConfig.ID

	return t.Suite.Equal(expected, actual, msgAndArgs)
}

func (t *RepositoryTest) NetEqual(expected model.NetConfig, actual model.NetConfig, msgAndArgs ...interface{}) bool {
	expected.CreatedAt = actual.CreatedAt
	expected.UpdatedAt = actual.UpdatedAt
	return t.Suite.Equal(expected, actual, msgAndArgs)
}

func (t *RepositoryTest) insertSampleData() {
	err := t.repository.db.Save(&room1).Error
	t.Require().NoError(err)
	err = t.repository.db.Save(&room2).Error
	t.Require().NoError(err)
	err = t.repository.db.Save(&room3).Error
	t.Require().NoError(err)
	err = t.repository.db.Save(&room4).Error
	t.Require().NoError(err)

	member1, err = t.repository.CreateOrUpdateMember(context.Background(), member1)
	t.Require().NoError(err)

	member2, err = t.repository.CreateOrUpdateMember(context.Background(), member2)
	t.Require().NoError(err)
}
