package netexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type IpAddrExpl struct{}

func (i *IpAddrExpl) RunExample(inputParams *goexpl.InputParams) error {
	// cidr := "1234:5678:abcd:ef12::/56"
	// cidr := "1234:5678:abcd:ffff::/59"
	// cidr := "1234:5678:abcd:ffff::/61"
	cidr := "2408:4008:5f3:8080::/60"
	if err := Ipv6ToSubnetDecimal(cidr); err != nil {
		return err
	}

	// testIpIndex := uint64(162274724273928191)
	// testIpIndex := uint64(162274724273928200)
	testIpIndex := uint64(71421922788834912)
	// prefixLength := 60
	prefixLength := 56
	longToCidrString(testIpIndex, prefixLength)

	return nil
}

func (i *IpAddrExpl) Init() {
	if err := goexpl.RegisterGoExample(i.GetGoExampleName(), &IpAddrExpl{}); err != nil {
		panic(err)
	}
}

func (i *IpAddrExpl) GetGoExampleName() string {
	return common.IpAddrExpl
}
