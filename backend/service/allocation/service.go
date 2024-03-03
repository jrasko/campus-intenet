package allocation

import (
	"backend/model"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"slices"
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
	// errNoIPsAllocatable if firstIP is >= broadcastIP
	errNoIPsAllocatable = errors.New("no allocatable ips")

	// errNoUnallocatedIP occurs if all possible ips are allocated
	errNoUnallocatedIP = model.Error(http.StatusInternalServerError,
		"no unallocated ip available", "no unallocated ip available")
)

func New(repo IPRepository, cidr string) Service {
	firstIP, broadcast, err := parseCIDR(cidr)
	if err != nil {
		panic(err)
	}
	log.Printf("[INFO] allocating IP in range %s to %s", firstIP, broadcast)

	return Service{
		repository: repo,

		broadcast: broadcast,
		firstIP:   firstIP,
	}
}

func (s Service) GetUnusedIP(ctx context.Context) (string, error) {
	ips, err := s.repository.GetAllIPs(ctx)
	if err != nil {
		return "", model.WrapGormError(err)
	}

	ip, err := s.getUnallocatedIP(ips)
	if err != nil {
		return "", fmt.Errorf("when allocating new ip: %w", err)
	}

	return ip, nil
}

func (s Service) getUnallocatedIP(allocatedIPs []string) (string, error) {
	for ip := s.firstIP; ip < s.broadcast; ip++ {
		if !slices.Contains(allocatedIPs, ip.String()) {
			return ip.String(), nil
		}
	}

	return "", errNoUnallocatedIP
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
		return 0, 0, errNoIPsAllocatable
	}

	return firstIP, broadcast, nil
}
