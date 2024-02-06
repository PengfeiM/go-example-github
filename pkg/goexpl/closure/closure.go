package closure

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Closure struct {
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func (c *Closure) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println("创建一个函数变量")
	nextInt := intSeq()

	/*
		nextInt() addr: 0x14000105578
		该地址前 8 字节存储了一个指针，addr: 0x1400005cef0
			接着 8 字节存储了一个指针，addr: 0x1400012e1a0
			接着 8 字节存储了一个指针，addr: 0x140001055f0

		* 0x1400005cef0 内容：
			- 0x00000001b384
	*/
	fmt.Println("调用 3 次")
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("创建一个函数变量")
	newInts := intSeq()
	fmt.Println(newInts())

	return nil
}

func (c *Closure) Init() {
	if err := goexpl.RegisterGoExample(c.GetGoExampleName(), &Closure{}); err != nil {
		panic(err)
	}
}

func (c *Closure) GetGoExampleName() string {
	return common.Closure
}
