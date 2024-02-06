// Package arrays
/*
在 golang 语言中，array 是一个特定长度的、编号的元素序列。
*/
package arrays

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type Arrays struct{}

func (a *Arrays) GetGoExampleName() string {
	return common.Arrays
}

func (a *Arrays) RunExample(inputParams *goexpl.InputParams) error {

	/*
		这里我们创建了一个 5 个元素的 array。
		array 比较重要的类型属性分别是它的元素类型和长度。
		在 golang 中，array 中的元素会默认赋零值。
		比如 int 类型的 array， 元素默认为0。
	*/
	var b [5]int
	fmt.Println("emp: ", b)

	b[4] = 100
	fmt.Println("set:", b)
	fmt.Println("get:", b[4])

	fmt.Println("len:", len(b))
	fmt.Println("cap:", cap(b))

	c := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	return nil
}

func (a *Arrays) Init() {
	if err := goexpl.RegisterGoExample(a.GetGoExampleName(), &Arrays{}); err != nil {
		panic(err)
	}
}
