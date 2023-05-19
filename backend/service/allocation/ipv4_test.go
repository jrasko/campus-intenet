package allocation

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
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
