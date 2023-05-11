package allocation

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findIPNotInList(t *testing.T) {
	t.Run("allocate in between", func(t *testing.T) {
		ips := []net.IP{
			net.IPv4(0, 0, 0, 5),
			net.IPv4(0, 0, 0, 6),
			net.IPv4(0, 0, 0, 7),
			net.IPv4(0, 0, 0, 8),
			net.IPv4(0, 0, 0, 10),
			net.IPv4(0, 0, 0, 11),
			net.IPv4(0, 0, 0, 12),
		}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(9), ip)
	})
	t.Run("empty list", func(t *testing.T) {
		ips := []net.IP{}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(5), ip)
	})
	t.Run("allocate at end", func(t *testing.T) {
		ips := []net.IP{
			net.IPv4(0, 0, 0, 5),
			net.IPv4(0, 0, 0, 6),
			net.IPv4(0, 0, 0, 7),
			net.IPv4(0, 0, 0, 8),
			net.IPv4(0, 0, 0, 9),
			net.IPv4(0, 0, 0, 10),
			net.IPv4(0, 0, 0, 11),
		}
		ip, err := findSuffixNotInList(ips, 5, 255)
		assert.NoError(t, err)
		assert.Equal(t, byte(12), ip)
	})
	t.Run("no space", func(t *testing.T) {
		ips := []net.IP{
			net.IPv4(0, 0, 0, 5),
			net.IPv4(0, 0, 0, 6),
			net.IPv4(0, 0, 0, 7),
			net.IPv4(0, 0, 0, 8),
			net.IPv4(0, 0, 0, 9),
			net.IPv4(0, 0, 0, 10),
			net.IPv4(0, 0, 0, 11),
		}
		ip, err := findSuffixNotInList(ips, 5, 11)
		assert.Error(t, err)
		assert.Equal(t, byte(255), ip)
	})
}
