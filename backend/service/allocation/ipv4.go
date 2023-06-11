package allocation

import (
	"encoding/binary"
	"net"
)

type IPv4 uint32

func NewIPv4(ip []byte) IPv4 {
	return IPv4(binary.BigEndian.Uint32(ip))
}

func (ip IPv4) String() string {
	ipBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(ipBytes, uint32(ip))
	return net.IP(ipBytes).To4().String()
}
