package allocation

import (
	"backend/model"
	"context"
	"errors"
	"net"
	"net/http"
)

//go:generate mockery --name IPRepository
type IPRepository interface {
	GetAllIPs(ctx context.Context) ([]string, error)
}

type Service struct {
	repository IPRepository

	broadcast IPv4
	firstIP   IPv4
}

var (
	// noIPsAllocatable if firstIP is >= broadcastIP
	noIPsAllocatable = errors.New("no allocatable ips")

	// noUnallocatedIP occurs if all possible ips are allocated
	noUnallocatedIP = model.Error(http.StatusInternalServerError,
		"no unallocated ip available", "no unallocated ip available")
)

func New(repo IPRepository, cidr string) Service {
	firstIP, broadcast, err := parseCIDR(cidr)
	if err != nil {
		panic(err)
	}

	return Service{
		repository: repo,

		broadcast: broadcast,
		firstIP:   firstIP,
	}
}

func (s Service) GetUnusedIP(ctx context.Context) (string, error) {
	ips, err := s.repository.GetAllIPs(ctx)
	if err != nil {
		return "", err
	}

	ip, err := s.getUnallocatedIP(ips)
	if err != nil {
		return "", err
	}

	return ip, nil
}

func (s Service) getUnallocatedIP(allocatedIPs []string) (string, error) {
	allocated := createAllocationMap(allocatedIPs)

	for ip := s.firstIP; ip < s.broadcast; ip++ {
		if !allocated[ip] {
			return ip.String(), nil
		}
	}

	return "", noUnallocatedIP
}

// createAllocationMap returns a maps IPv4 -> bool where every IP on allocatedIPs is mapped to true
// the trick behind that is, that one can simply query the map and get true, if the ip is already allocated
// or else false
func createAllocationMap(allocatedIPs []string) map[IPv4]bool {
	m := map[IPv4]bool{}
	for _, ipString := range allocatedIPs {
		ip := NewIPv4(net.ParseIP(ipString).To4())
		m[ip] = true
	}
	return m
}

// parseCIDR parses a cidr string like 192.168.1.4/24
// and returns the first allocatable IPv4 and the broadcast IPv4
// A successful parseCIDR returns err == nil, else err contains the error
func parseCIDR(cidr string) (firstIP IPv4, broadcast IPv4, err error) {
	firstIPRaw, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return 0, 0, err
	}

	firstIP = NewIPv4(firstIPRaw.To4())
	netIP := NewIPv4(ipNet.IP)
	mask := NewIPv4(ipNet.Mask)

	broadcast = ^mask | netIP

	// check if there are allocatable ips
	if firstIP >= broadcast {
		return 0, 0, noIPsAllocatable
	}

	return firstIP, broadcast, nil
}
