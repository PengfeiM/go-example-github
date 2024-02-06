// Package interfaceexpl
/*
Interface 接口是一系列方法声明的集合
*/

package interfaceexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"math"
)

// 实现接口 goexpl.GoExample，用于接口类型的示例
type InterfaceExpl struct{}

type geometry interface {
	area() float64
	perim() float64
}

// 以下为 geometry 这个接口的两种实现。分别是 长方形 rect，圆 circle。
// 它们都实现了方法 area() 和 perim()。
// 同时，它们定义了自己的字段。

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.height + 2*r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return math.Pi * c.radius * 2
}

// measure
// 定义方法 measure()，入参为 geometry，输入其实现 rect 和 circle 均可。
// 输出对应的 area 和 perim
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func (i *InterfaceExpl) RunExample(inputParams *goexpl.InputParams) error {
	r := rect{4, 3}
	c := circle{5}

	measure(r)
	measure(c)

	return nil
}

func (i *InterfaceExpl) Init() {
	if err := goexpl.RegisterGoExample(i.GetGoExampleName(), &InterfaceExpl{}); err != nil {
		panic(err)
	}
}

func (i *InterfaceExpl) GetGoExampleName() string {
	return common.InterfaceExpl
}
