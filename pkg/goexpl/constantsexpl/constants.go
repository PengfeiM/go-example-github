package constantsexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"math"
)

type Constants struct {
}

func (c *Constants) GetGoExampleName() string {
	return common.Constants
}

const s string = "constant"

func (c *Constants) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
	return nil
}

func (c *Constants) Init() {
	if err := goexpl.RegisterGoExample(c.GetGoExampleName(), &Constants{}); err != nil {
		panic(err)
	}
}
