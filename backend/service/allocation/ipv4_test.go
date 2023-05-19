package allocation

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
)

func Test_findIPNotInList(t *testing.T) {
	t.Run("allocate in between", func(t *testing.T) {
		ips := []string{
			net.IPv4(0, 0, 0, 5).String(),
			net.IPv4(0, 0, 0, 6).String(),
			net.IPv4(0, 0, 0, 7).String(),
			net.IPv4(0, 0, 0, 8).String(),
			net.IPv4(0, 0, 0, 10).String(),
			net.IPv4(0, 0, 0, 11).String(),
			net.IPv4(0, 0, 0, 12).String(),
		}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(9), ip)
	})
	t.Run("empty list", func(t *testing.T) {
		ips := []string{}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(5), ip)
	})
	t.Run("allocate at end", func(t *testing.T) {
		ips := []string{
			net.IPv4(0, 0, 0, 5).String(),
			net.IPv4(0, 0, 0, 6).String(),
			net.IPv4(0, 0, 0, 7).String(),
			net.IPv4(0, 0, 0, 8).String(),
			net.IPv4(0, 0, 0, 9).String(),
			net.IPv4(0, 0, 0, 10).String(),
			net.IPv4(0, 0, 0, 11).String(),
		}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(12), ip)
	})
	t.Run("no space", func(t *testing.T) {
		ips := []string{
			net.IPv4(0, 0, 0, 5).String(),
			net.IPv4(0, 0, 0, 6).String(),
			net.IPv4(0, 0, 0, 7).String(),
			net.IPv4(0, 0, 0, 8).String(),
			net.IPv4(0, 0, 0, 9).String(),
			net.IPv4(0, 0, 0, 10).String(),
			net.IPv4(0, 0, 0, 11).String(),
		}
		ip, err := findSuffixNotInList(ips, 5, 11)
		assert.Error(t, err)
		assert.Equal(t, byte(255), ip)
	})
}

func TestService_GetUnusedIP(t *testing.T) {
	repo := NewMockIPRepository(t)

	t.Run("it works in default case", func(t *testing.T) {
		repo.EXPECT().GetAllIPs(ctx).Return([]string{}, nil).Once()

		ipService := New(repo)
		ip, err := ipService.GetUnusedIP(ctx)
		assert.NoError(t, err)
		assert.Equal(t, "149.201.243.6", ip)
	})
	t.Run("it returns an error on db missbehavior", func(t *testing.T) {
		anError := errors.New("error")
		repo.EXPECT().
			GetAllIPs(ctx).
			Return([]string{"test"}, anError).
			Once()

		ipService := New(repo)
		ip, err := ipService.GetUnusedIP(ctx)
		assert.ErrorIs(t, err, anError)
		assert.Equal(t, "", ip)
	})
	t.Run("it returns an error if no unallocated ip was found", func(t *testing.T) {
		repo.EXPECT().
			GetAllIPs(ctx).
			Return([]string{"149.201.243.6", "149.201.243.7", "149.201.243.8"}, nil).
			Once()

		ipService := New(repo)
		ipService.maxSuffix = 8

		ip, err := ipService.GetUnusedIP(ctx)
		assert.ErrorIs(t, err, noIPFoundErr)
		assert.Equal(t, "", ip)
	})
}
