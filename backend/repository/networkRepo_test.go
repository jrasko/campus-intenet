package repository

import (
	"backend/model"
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ctx = context.Background()
)

func setupDB() (MemberRepository, error) {
	cfg := model.LoadConfig()
	repo, err := New(cfg.DSN())
	if err != nil {
		return MemberRepository{}, err
	}
	repo.db.Exec("TRUNCATE member_configs RESTART IDENTITY")
	return repo, err
}

func TestNew(t *testing.T) {

	t.Run("empty dsn", func(t *testing.T) {
		repo, err := New("")
		assert.Error(t, err)
		assert.Equal(t, MemberRepository{}, repo)
	})
	t.Run("it creates table on setup", func(t *testing.T) {
		config := model.LoadConfig()
		repo, err := New(config.DSN())
		assert.NoError(t, err)
		repo.db.Exec(fmt.Sprintf("DROP TABLE %s", memberTable))

		_, err = New(config.DSN())
		assert.NoError(t, err)
	})
}

func TestMemberRepository(t *testing.T) {
	repo, creationErr := setupDB()
	require.NoError(t, creationErr)

	member := model.MemberConfig{
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
	member2 := model.MemberConfig{
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
	updatedMember := model.MemberConfig{
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
		newMember, err := repo.UpdateMemberConfig(ctx, member)
		assert.NoError(t, err)
		assert.NotEmpty(t, newMember.ID)
		member.ID = newMember.ID
		assert.Equal(t, member, newMember)
	})
	t.Run("it creates another member", func(t *testing.T) {
		newMember, err := repo.UpdateMemberConfig(ctx, member2)
		member2.ID = newMember.ID
		assert.NoError(t, err)
		assert.Equal(t, member2, newMember)
	})
	t.Run("it checks unique constraints", func(t *testing.T) {
		newMember := model.MemberConfig{
			Mac:    member.Mac,
			RoomNr: member.RoomNr,
			IP:     member.IP,
		}
		newMember, err := repo.UpdateMemberConfig(ctx, newMember)
		assert.Equal(t, http.StatusConflict, err.(model.HttpError).Status())
	})
	t.Run("it retrevies a single member", func(t *testing.T) {
		m, err := repo.GetMemberConfig(ctx, member.ID)
		assert.NoError(t, err)
		assert.Equal(t, m, m)
	})
	t.Run("it retreives multiple members", func(t *testing.T) {
		members, err := repo.GetAllMemberConfigs(ctx)
		assert.NoError(t, err)
		assert.Len(t, members, 2)
		assert.Contains(t, members, member)
		assert.Contains(t, members, member2)
	})
	t.Run("it retreives all ips", func(t *testing.T) {
		ips, err := repo.GetAllIPs(ctx)
		assert.NoError(t, err)
		assert.Len(t, ips, 2)
		assert.Contains(t, ips, member.IP)
		assert.Contains(t, ips, member2.IP)
	})
	t.Run("it retreives all macs", func(t *testing.T) {
		macs, err := repo.GetAllMacs(ctx)
		assert.NoError(t, err)
		assert.Len(t, macs, 2)
		assert.Contains(t, macs, member.Mac)
		assert.Contains(t, macs, member2.Mac)
	})
	t.Run("it updates a member", func(t *testing.T) {
		updatedMember.ID = member.ID
		newMember, err := repo.UpdateMemberConfig(ctx, updatedMember)
		assert.NoError(t, err)
		assert.Equal(t, updatedMember, newMember)
	})
	t.Run("it retreives first and lastnames", func(t *testing.T) {
		persons, err := repo.GetNonPayingMembers(ctx)
		assert.NoError(t, err)
		assert.Len(t, persons, 1)
		assert.Contains(t, persons,
			model.MemberConfig{
				Firstname: member2.Firstname,
				Lastname:  member2.Lastname,
			},
		)

	})
	t.Run("it resets payments", func(t *testing.T) {
		err := repo.ResetPayment(ctx)
		assert.NoError(t, err)

		members, err := repo.GetAllMemberConfigs(ctx)
		assert.NoError(t, err)
		for _, m := range members {
			assert.False(t, m.HasPaid)
		}
	})
	t.Run("it deletes a member", func(t *testing.T) {
		err := repo.DeleteMemberConfig(ctx, member2.ID)
		assert.NoError(t, err)

		_, err = repo.GetMemberConfig(ctx, member2.ID)
		assert.Equal(t, http.StatusNotFound, err.(model.HttpError).Status())
	})

}

func TestManyMembers(t *testing.T) {
	t.SkipNow()
	repo, creationErr := setupDB()
	require.NoError(t, creationErr)
	for i := 0; i < 100; i++ {
		member := model.MemberConfig{
			Firstname: strconv.Itoa(i),
			Lastname:  strconv.Itoa(i),
			Mac:       "00:00:00:00:00:" + hex.EncodeToString([]byte{byte(i)}),
			RoomNr:    strconv.Itoa(i),
			IP:        strconv.Itoa(i),
		}
		member, err := repo.UpdateMemberConfig(ctx, member)
		require.NoError(t, err)
	}
}
