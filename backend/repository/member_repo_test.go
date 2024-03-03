package repository

import (
	"backend/model"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	ctx = context.Background()
)

func setupDB() (MemberRepository, error) {
	cfg, err := model.LoadConfig(context.Background())
	if err != nil {
		return MemberRepository{}, err
	}
	cfg.DBDatabase = "testing"
	repo, err := New(cfg.DSN())
	if err != nil {
		return MemberRepository{}, err
	}
	repo.db.Exec("TRUNCATE member_configs RESTART IDENTITY")
	return repo, nil
}

func TestNew(t *testing.T) {
	t.Run("empty dsn", func(t *testing.T) {
		repo, err := New("")
		assert.Error(t, err)
		assert.Equal(t, MemberRepository{}, repo)
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

var (
	member = model.Member{
		Firstname: "first",
		Lastname:  "name",
		NetConfig: model.NetConfig{
			Mac: "00:11:22:33:44:55",
			IP:  "192.168.1.1",
		},
		RoomNr: "00-11",
		Room: model.Room{
			Number: "00-11",
			WG:     "25a",
		},
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
		Room: model.Room{
			Number: "99-99",
			WG:     "13x",
		},
		Email: "mail@email.email",
		Phone: "0999888777666",
	}
	updatedMember = model.Member{
		Firstname: "Bernd",
		Lastname:  "Das Brot",
		NetConfig: model.NetConfig{
			Mac: "55:44:33:22:11:00",
			IP:  "172.0.0.1",
		},
		HasPaid: true,
		RoomNr:  "11-00",
		Room: model.Room{
			Number: "11-00",
			WG:     "29b",
		},
		Email: "other-email@alumni.test-provider.com",
		Phone: "987654321",
	}
)

func TestMemberRepository(t *testing.T) {
	repo, creationErr := setupDB()
	require.NoError(t, creationErr)

	t.Run("it creates a member", func(t *testing.T) {
		newMember, err := repo.CreateOrUpdateMember(ctx, member)
		assert.NoError(t, err)
		assert.NotEmpty(t, newMember.ID)
		member.ID = newMember.ID

		overwrite(t, &newMember, &member)
		assert.Equal(t, member, newMember)
	})
	t.Run("it creates another member", func(t *testing.T) {
		newMember, err := repo.CreateOrUpdateMember(ctx, member2)
		assert.NoError(t, err)
		member2 = newMember
	})
	t.Run("it retrevies a single member", func(t *testing.T) {
		m, err := repo.GetMember(ctx, member.ID)
		assert.NoError(t, err)
		overwrite(t, &m, &member)
		assert.Equal(t, member, m)
	})
	t.Run("it retreives multiple members", func(t *testing.T) {
		members, err := repo.ListMembers(ctx, model.MemberRequestParams{})
		assert.NoError(t, err)
		assert.Len(t, members, 3)
		assert.Equal(t, members[0].ID, member.ID)
		assert.Equal(t, members[2].ID, member2.ID)
	})
	t.Run("it searches for members", func(t *testing.T) {
		members, err := repo.ListMembers(ctx, model.MemberRequestParams{Search: "first"})
		assert.NoError(t, err)
		assert.Len(t, members, 1)
		assert.Equal(t, members[0].ID, member.ID)
	})

	t.Run("it updates a member", func(t *testing.T) {
		updatedMember.ID = member.ID
		newMember, err := repo.CreateOrUpdateMember(ctx, updatedMember)
		assert.NoError(t, err)
		assert.Equal(t, updatedMember.Firstname, newMember.Firstname)
	})

	t.Run("it resets payments", func(t *testing.T) {
		err := repo.ResetPayment(ctx)
		assert.NoError(t, err)

		members, err := repo.ListMembers(ctx, model.MemberRequestParams{})
		assert.NoError(t, err)
		for _, m := range members {
			assert.False(t, m.HasPaid)
		}
	})
	t.Run("it deletes a member", func(t *testing.T) {
		err := repo.DeleteMembers(ctx, member2.ID)
		assert.NoError(t, err)

		_, err = repo.GetMember(ctx, member2.ID)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func overwrite(t *testing.T, fromDB *model.Member, compare *model.Member) {
	assert.NotEmpty(t, fromDB.NetConfigID)
	assert.NotEmpty(t, fromDB.NetConfig.ID)
	assert.NotEmpty(t, fromDB.RoomNr)
	assert.NotEmpty(t, fromDB.Room.Number)

	assert.NotEmpty(t, fromDB.CreatedAt)
	assert.NotEmpty(t, fromDB.UpdatedAt)
	fromDB.CreatedAt = compare.CreatedAt
	fromDB.UpdatedAt = compare.UpdatedAt

	fromDB.NetConfigID = compare.NetConfigID
	fromDB.NetConfig.ID = compare.NetConfig.ID
	fromDB.NetConfig.CreatedAt = compare.NetConfig.CreatedAt
	fromDB.NetConfig.UpdatedAt = compare.NetConfig.UpdatedAt
}
