package forexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type ForExpl struct {
}

func (f *ForExpl) GetGoExampleName() string {
	return common.ForExpl
}

func (f *ForExpl) RunExample(inputParams *goexpl.InputParams) error {
	var i int
	i = 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	var n int
	for n = 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	return nil
}

func (f *ForExpl) Init() {
	if err := goexpl.RegisterGoExample(f.GetGoExampleName(), &ForExpl{}); err != nil {
		panic(err)
	}
}
