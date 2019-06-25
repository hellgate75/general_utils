package net

import (
	"errors"
	"fmt"
	"net"
)

type IPAddressInfo struct {
	Address            string
	Ip                 net.IP
	Mask               net.IP
	Loopback           bool
	GlobalUnicast      bool
	LocalUnicast       bool
	MultiCast          bool
	IfcLocalMulticast  bool
	LinkLocalMulticast bool
}

// Easiest way to create net.IP value is to use
// net.ParseIP which parses a string value representation
// of a IPv4 dot-separated or IPv6 colon-separated address.
// This example uses net.ParseIP to parse an IP address provided
// from the command line and prints information about the address.
func GetIpAddressInfo(ipAddr string) (*IPAddressInfo, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return nil, errors.New(fmt.Sprintf("Unable to parse IP address : %s. Address should use IPv4 dot-notation or IPv6 colon-notation.", ipAddr))
	}

	info := IPAddressInfo{
		Address:            ipAddr,
		Ip:                 ip,
		Mask:               net.IP(ip.DefaultMask()),
		Loopback:           ip.IsLoopback(),
		GlobalUnicast:      ip.IsGlobalUnicast(),
		LocalUnicast:       ip.IsLinkLocalUnicast(),
		MultiCast:          ip.IsMulticast(),
		IfcLocalMulticast:  ip.IsInterfaceLocalMulticast(),
		LinkLocalMulticast: ip.IsLinkLocalMulticast(),
	}
	return &info, nil
}
