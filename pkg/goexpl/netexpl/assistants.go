package netexpl

import (
	"fmt"
	"net"
	"strconv"
)

func Ipv6ToSubnetDecimal(cidr string) error {
	ip, net, err := net.ParseCIDR(cidr)
	if err != nil {
		return fmt.Errorf("invalid Ipv6 cidr: %s", cidr)
	}

	// first way to calculate ipIndex1
	ipIndex1 := (uint64(
		ip[0],
	) << 48) + (uint64(
		ip[1],
	) << 40) + (uint64(
		ip[2],
	) << 32) + (uint64(
		ip[3],
	) << 24) + (uint64(
		ip[4],
	) << 16) + (uint64(
		ip[5],
	) << 8) + uint64(
		ip[6],
	)
	fmt.Printf("first way to calculate ipIndex: %d\n", ipIndex1)

	// my way
	// Extract the first 64 bits (network identifier part of IPv6)
	upperPart := uint64(0)
	for i := 0; i < 7; i++ { // to make it same as biz code, use 7 here
		upperPart = (upperPart << 8) | uint64(ip[i])
	}
	prefixLength, _ := net.Mask.Size()
	// mask := uint64((1<<prefixLength) -1 )
	ipIndex2 := (((upperPart << 8) | uint64(ip[7])) >> (64 - prefixLength))
	fmt.Printf("my way to calculate ipIndex: %d\n", ipIndex2)

	fmt.Printf("is to ipIndex equal? %t\n", ipIndex1 == ipIndex2)

	fmt.Printf("diff between 2 ways: %d\n", ipIndex2-(1<<(prefixLength-56))*ipIndex1)

	return nil
}

func longToCidrString(i uint64, prefixLength int) error {
	// first
	ip := make(net.IP, net.IPv6len)
	ip[0] = byte(i >> 48)
	ip[1] = byte((i >> 40) & 0xff)
	ip[2] = byte((i >> 32) & 0xff)
	ip[3] = byte((i >> 24) & 0xff)
	ip[4] = byte((i >> 16) & 0xff)
	ip[5] = byte((i >> 8) & 0xff)
	ip[6] = byte(i & 0xff)

	cidrBlock := ip.String() + "/56"

	fmt.Printf("old way: %s\n", cidrBlock)

	// second, with flexible prefix
	// process last section
	ip2 := make(net.IP, net.IPv6len)
	leftShift := 64 - prefixLength
	tmpIndex := i << leftShift
	ip2[0] = byte(tmpIndex >> 56)
	ip2[1] = byte((tmpIndex >> 48) & 0xff)
	ip2[2] = byte((tmpIndex >> 40) & 0xff)
	ip2[3] = byte((tmpIndex >> 32) & 0xff)
	ip2[4] = byte((tmpIndex >> 24) & 0xff)
	ip2[5] = byte((tmpIndex >> 16) & 0xff)
	ip2[6] = byte((tmpIndex >> 8) & 0xff)
	ip2[7] = byte(tmpIndex & 0xff)

	cidrBlock2 := ip2.String() + "/" + strconv.Itoa(prefixLength)
	fmt.Printf("new way: %s\n", cidrBlock2)

	return nil
}
