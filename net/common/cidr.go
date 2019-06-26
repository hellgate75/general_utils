package common

import (
	"errors"
	"fmt"
	"math"
	"net"
)

type CIDRInfo struct {
	Cidr       string
	Ip         net.IP
	LastIp     net.IP
	IpNet      net.IPNet
	Size       int
	TotalHosts float64
	WildcardIP net.IP
}

// This function implements a CIDR subnet calculator.
// It takes a CIDR address prefix an calculates ip-ranges,
// total hosts, wildcard mask, etc. all packed in a type CIDRInfo struct instance
func GetCIdrInfo(cidr string) (*CIDRInfo, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Println("failed parsing CIDR address: ", err)
		return nil, errors.New(fmt.Sprintf("Failed parsing CIDR address: %s, cause: %s", cidr, err.Error()))
	}

	// Given IPv4 block 192.168.100.14/24
	// The followings uses IPNet to get:
	// - The routing address for the subnet (i.e. 192.168.100.0)
	// - one-bits of the network mask (24 out of 32 total)
	// - The subnetmask (i.e. 255.255.255.0)
	// - Total hosts in the network (2 ^(host identifer bits) or 2^8)
	// - Wildcard the inverse of subnet mask (i.e. 0.0.0.255)
	// - The maximum address of the subnet (i.e. 192.168.100.255)
	ones, totalBits := ipnet.Mask.Size()
	size := totalBits - ones                 // usable bits
	totalHosts := math.Pow(2, float64(size)) // 2^size
	wildcardIP := __wildcard(net.IP(ipnet.Mask))
	last := __lastIP(ip, net.IPMask(wildcardIP))
	info := CIDRInfo{
		Cidr:       cidr,
		Ip:         ip,
		LastIp:     last,
		IpNet:      *ipnet,
		Size:       size,
		TotalHosts: totalHosts,
		WildcardIP: wildcardIP,
	}
	return &info, nil
}

// wildcard returns the opposite of the
// the netmask for the network.
func __wildcard(mask net.IP) net.IP {
	var ipVal net.IP
	for _, octet := range mask {
		ipVal = append(ipVal, ^octet)
	}
	return ipVal
}

// lastIP calculates the highest addressable IP for given
// for a given subnet. It Loops through each octet of the
// subnet's IP address and applies bitwise OR operation
// to each corresponding octet from the mask value.
func __lastIP(ip net.IP, mask net.IPMask) net.IP {
	ipIn := ip.To4() // is it an IPv4
	if ipIn == nil {
		ipIn = ip.To16() // is it IPv6
		if ipIn == nil {
			return nil
		}
	}
	var ipVal net.IP
	// apply network mask to each octet
	for i, octet := range ipIn {
		ipVal = append(ipVal, octet|mask[i])
	}
	return ipVal
}
