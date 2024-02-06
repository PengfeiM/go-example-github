package recursion

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"strconv"
)

type Recursion struct {
}

func fact(n int) int {
	if n < 0 {
		panic("illegal input of num")
	}
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func (r *Recursion) RunExample(inputParams *goexpl.InputParams) error {
	argc := inputParams.GetArgNum()
	var n int
	if argc == 0 {
		n = 7
	} else if argc > 1 {
		panic("too much input params")
	} else {
		n, _ = strconv.Atoi(inputParams.GetArgs()[0])
	}

	fmt.Println(fact(n))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))

	return nil
}

func (r *Recursion) Init() {
	if err := goexpl.RegisterGoExample(r.GetGoExampleName(), &Recursion{}); err != nil {
		panic(err)
	}
}

func (r *Recursion) GetGoExampleName() string {
	return common.Recursion
}
