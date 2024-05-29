package shiftexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"strconv"
)

type ShiftExpl struct{}

func (s *ShiftExpl) RunExample(inputParams *goexpl.InputParams) error {
	var i int
	var len int

	if inputParams.GetArgNum() < 1 {
		len = 17
		fmt.Println("no input mask len, use default value")
	} else {
		len, _ = strconv.Atoi(inputParams.GetArgs()[0])
	}

	i = 1
	fmt.Println("i =", i, "len =", len)
	fmt.Println("fmt.Println(i << (32 - len))")
	fmt.Println(i << (32 - len))

	return nil
}

func (s *ShiftExpl) Init() {
	err := goexpl.RegisterGoExample(s.GetGoExampleName(), &ShiftExpl{})
	if err != nil {
		panic(err)
	}
}

func (s *ShiftExpl) GetGoExampleName() string {
	return common.ShiftExpl
}
