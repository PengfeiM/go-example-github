package netexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"net"
)

type NetExpl struct{}

func (n *NetExpl) RunExample(inputParams *goexpl.InputParams) error {
	cidrBlock := "192.168.0.0/24"

	ip, ipNet, _ := net.ParseCIDR(cidrBlock)
	fmt.Println(ip)
	fmt.Println(ipNet)

	fmt.Println(ipNet.Mask)
	fmt.Println(ipNet.Mask.Size())
	fmt.Println(ipNet.Mask.String())

	return nil
}

func (n *NetExpl) Init() {
	err := goexpl.RegisterGoExample(n.GetGoExampleName(), &NetExpl{})
	if err != nil {
		panic(err)
	}
}

func (n *NetExpl) GetGoExampleName() string {
	return common.NetExpl
}
