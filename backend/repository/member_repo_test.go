package repository

import (
	"backend/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestNew(t *testing.T) {
	t.Run("empty dsn", func(t *testing.T) {
		repo, err := New("")
		assert.Error(t, err)
		assert.Equal(t, Repository{}, repo)
	})
	t.Run("it creates table on setup", func(t *testing.T) {
		config, err := model.LoadConfig(context.Background())
		require.NoError(t, err)
		config.DBDatabase = "testing"
		repo, err := New(config.DSN())
		assert.NoError(t, err)
		err = repo.db.Migrator().DropTable(&model.Member{}, &model.Room{}, &model.NetConfig{})
		require.NoError(t, err)

		_, err = New(config.DSN())
		assert.NoError(t, err)
	})
}

func (t *RepositoryTest) TestMemberRepository_CreateOrUpdateMember_create() {
	newMember := model.Member{
		Firstname: "firstname",
		Lastname:  "lastname",
		HasPaid:   true,
		RoomNr:    room3.Number,
		Room:      room3,
		NetConfig: model.NetConfig{
			Mac:          "AA:BB:CC:DD:EE:FF",
			IP:           "192.168.0.1",
			Manufacturer: "test",
			Disabled:     true,
		},
	}
	fromDB, err := t.repository.CreateOrUpdateMember(context.Background(), newMember)
	t.NoError(err)
	t.NotEmpty(fromDB.ID)
	newMember.ID = fromDB.ID
	newMember.NetConfigID = fromDB.NetConfigID

	t.MemberEqual(newMember, fromDB)
}

func (t *RepositoryTest) TestMemberRepository_CreateOrUpdateMember_create_requireExistingRoom() {
	newMember := model.Member{
		Firstname: "firstname",
		Lastname:  "lastname",
		RoomNr:    "new",
		NetConfig: model.NetConfig{
			Mac: "00:00:00:11:11:11",
			IP:  "192.168.1.1",
		},
	}
	_, err := t.repository.CreateOrUpdateMember(context.Background(), newMember)
	t.Error(err)
}

func (t *RepositoryTest) TestMemberRepository_CreateOrUpdateMember_update() {
	updatedMember := model.Member{
		ID:          member1.ID,
		Firstname:   "Bernd",
		Lastname:    "Das Brot",
		NetConfigID: member1.NetConfigID,
		NetConfig: model.NetConfig{
			ID:           member1.NetConfigID,
			Mac:          "55:44:33:22:11:00",
			IP:           "172.0.0.1",
			Manufacturer: "other",
		},
		HasPaid: true,
		RoomNr:  room4.Number,
		Room:    room4,
		Email:   "other-email@alumni.test-provider.com",
		Phone:   "987654321",
	}
	fromDB, err := t.repository.CreateOrUpdateMember(context.Background(), updatedMember)
	t.NoError(err)
	t.MemberEqual(updatedMember, fromDB)
}

func (t *RepositoryTest) TestMemberRepository_GetMember() {
	m, err := t.repository.GetMember(context.Background(), member1.ID)
	t.NoError(err)
	t.MemberEqual(member1, m)
}

func (t *RepositoryTest) TestMemberRepository_DeleteMember() {
	err := t.repository.DeleteMember(context.Background(), member2.ID)
	t.NoError(err)

	_, err = t.repository.GetMember(context.Background(), member2.ID)
	t.Equal(gorm.ErrRecordNotFound, err)
}
