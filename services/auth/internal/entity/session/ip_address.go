package session

import "net"

// IPAddress represents a validated IP address string.
type IPAddress string

func NewIPAddress(value string) (IPAddress, error) {
	ip := IPAddress(value)
	if err := ip.Validate(); err != nil {
		return "", err
	}
	return ip, nil
}

// Validate checks whether the IPAddress is non-empty and valid.
func (ip IPAddress) Validate() error {
	if ip == "" {
		return ErrEmptyIpAddress
	}
	parsed := net.ParseIP(ip.String())
	if parsed == nil {
		return ErrInvalidIpAddress
	}
	return nil
}

// String returns the string representation of the IPAddress.
func (ip IPAddress) String() string {
	return string(ip)
}
