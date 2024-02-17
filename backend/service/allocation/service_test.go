package allocation

import (
	"backend/model"
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func Test_New(t *testing.T) {
	t.Run("standard create", func(t *testing.T) {
		New(NewMockIPRepository(t), "192.128.2.5/28")
	})
	t.Run("too few allocatables", func(t *testing.T) {
		assert.Panics(t, func() {
			New(NewMockIPRepository(t), "192.168.2.255/24")
		})
	})
	t.Run("invalid cidr", func(t *testing.T) {
		assert.Panics(t, func() {
			_ = New(NewMockIPRepository(t), "192.168.2.2")
		})
	})

}

func Test_getUnallocatedIP(t *testing.T) {
	s := New(NewMockIPRepository(t), "0.0.0.2/29")

	t.Run("allocate in between", func(t *testing.T) {
		ips := []string{
			"0.0.0.2",
			"0.0.0.3",
			"0.0.0.4",
			"0.0.0.6",
		}
		ip, err := s.getUnallocatedIP(ips)
		assert.NoError(t, err)
		assert.Equal(t, "0.0.0.5", ip)
	})
	t.Run("empty list", func(t *testing.T) {
		var ips []string
		ip, err := s.getUnallocatedIP(ips)
		assert.NoError(t, err)
		assert.Equal(t, "0.0.0.2", ip)
	})
	t.Run("allocate at end", func(t *testing.T) {
		ips := []string{
			"0.0.0.2",
			"0.0.0.3",
			"0.0.0.4",
			"0.0.0.5",
		}
		ip, err := s.getUnallocatedIP(ips)
		assert.NoError(t, err)
		assert.Equal(t, "0.0.0.6", ip)
	})
	t.Run("no space", func(t *testing.T) {
		ips := []string{
			"0.0.0.2",
			"0.0.0.3",
			"0.0.0.4",
			"0.0.0.5",
			"0.0.0.6",
		}
		ip, err := s.getUnallocatedIP(ips)
		assert.Error(t, err)
		assert.Equal(t, "", ip)
	})
}

func TestService_GetUnusedIP(t *testing.T) {
	repo := NewMockIPRepository(t)
	ipService := New(repo, "149.201.243.6/24")

	t.Run("it works in default case", func(t *testing.T) {
		repo.EXPECT().
			GetAllIPs(ctx).
			Return([]string{}, nil).
			Once()

		ip, err := ipService.GetUnusedIP(ctx)
		assert.NoError(t, err)
		assert.Equal(t, "149.201.243.6", ip)
	})
	t.Run("it returns an error on db misbehavior", func(t *testing.T) {
		anError := errors.New("error")
		repo.EXPECT().
			GetAllIPs(ctx).
			Return([]string{"test"}, anError).
			Once()

		ip, err := ipService.GetUnusedIP(ctx)
		assert.Equal(t, http.StatusInternalServerError, err.(model.HttpError).Status())
		assert.Equal(t, "", ip)
	})
	t.Run("it returns an error if no unallocated ip was found", func(t *testing.T) {
		repo.EXPECT().
			GetAllIPs(ctx).
			Return([]string{"149.201.243.5", "149.201.243.6"}, nil).
			Once()

		ipService = New(repo, "149.201.243.5/29")

		ip, err := ipService.GetUnusedIP(ctx)
		assert.Equal(t, "", ip)
		assert.ErrorIs(t, err, errNoUnallocatedIP)
	})
}
