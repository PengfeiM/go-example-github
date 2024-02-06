package ifel

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type IfEl struct {
}

func (i *IfEl) GetGoExampleName() string {
	return common.IfEl
}

func (i *IfEl) RunExample(inputParams *goexpl.InputParams) error {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	} else {
		fmt.Println("8 is not divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit, and it's positive")
	} else {
		fmt.Println(num, "has multiple digits, and it's positive")
	}

	return nil
}

func (i *IfEl) Init() {
	if err := goexpl.RegisterGoExample(i.GetGoExampleName(), &IfEl{}); err != nil {
		panic(err)
	}
}
