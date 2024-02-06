// Package pointer
/*
	golang 支持指针，可以用于传递变量的引用
*/
package pointer

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Pointer struct {
}

// 这个函数的参数是传值，函数内部是复制了一份变量。
// 函数内部的改变不会影响外部的变量
func zeroval(ival int) {
	ival = 0
}

// 这个函数的参数是指针，可以传引用
// 在这里修改指针指向的变量值，在外部访问这个指针指向的变量时会发现也发生变化
func zeroptr(iptr *int) {
	*iptr = 0
}

func (p *Pointer) RunExample(inputParams *goexpl.InputParams) error {
	i := 1
	fmt.Println("initial:", i)
	// initial: 1

	zeroval(i)
	fmt.Println("zeroval:", i)
	// zeroval: 1

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
	// zeroptr: 1

	/*
		zeroval() 不会更改 main 中的 i，但 zeroptr() 会更改，因为它具有对该变量的内存地址的引用。
	*/

	fmt.Println("pointer:", &i)
	fmt.Printf("pointer: %p\n", &i)

	return nil
}

func (p *Pointer) Init() {
	if err := goexpl.RegisterGoExample(p.GetGoExampleName(), &Pointer{}); err != nil {
		panic(err)
	}
}

func (p *Pointer) GetGoExampleName() string {
	return common.Pointer
}
