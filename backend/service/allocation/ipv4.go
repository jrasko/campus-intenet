package allocation

import (
	"backend/model"
	"context"
	"net"
	"net/http"
)

//go:generate mockery --name IPRepository
type IPRepository interface {
	GetAllIPs(ctx context.Context) ([]string, error)
}

type Service struct {
	repository IPRepository
	minSuffix  byte
	maxSuffix  byte
	IPPrefix   []byte
}

func New(repo IPRepository) Service {
	return Service{
		repository: repo,
		minSuffix:  byte(6),
		maxSuffix:  byte(254),
		IPPrefix:   []byte{149, 201, 243},
	}
}

func (s Service) GetUnusedIP(ctx context.Context) (string, error) {
	ips, err := s.repository.GetAllIPs(ctx)
	if err != nil {
		return "", err
	}

	suffix, err := findSuffixNotInList(ips, s.minSuffix, s.maxSuffix)
	if err != nil {
		return "", err
	}

	return net.IPv4(s.IPPrefix[0], s.IPPrefix[1], s.IPPrefix[2], suffix).String(), nil
}

func findSuffixNotInList(ips []string, minSuffix byte, maxSuffix byte) (byte, error) {
	for i := minSuffix; i <= maxSuffix; i++ {
		listIndex := i - minSuffix
		if listIndex >= byte(len(ips)) {
			return i, nil
		}
		if address := []byte(net.ParseIP(ips[listIndex])); address[15] != i {
			return i, nil
		}
	}
	return 0, model.Error(http.StatusInternalServerError, "no unused ips available", "no unallocated ip available")

}