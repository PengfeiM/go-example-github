// Package methodexpl
// 结构体可以定义依赖于结构体的方法 method

package methodexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"reflect"
)

// MethodExpl
// 结构体方法示例
type MethodExpl struct {
}

// rect
// 长方形结构体，用于测试结构体方法
type rect struct {
	Width, Height int
}

func (r *rect) area() int {
	return r.Width * r.Height
}

func (r rect) perim() int {
	return 2*r.Height + 2*r.Width
}

func (r rect) setWithValue(field string, value int) {
	v := reflect.ValueOf(&r)

	v.Elem().FieldByName(field).SetInt(int64(value))
}

func (r *rect) setWithPtr(field string, value int) {
	v := reflect.ValueOf(r)

	v.Elem().FieldByName(field).SetInt(int64(value))
}

func (m *MethodExpl) RunExample(inputParams *goexpl.InputParams) error {
	r := rect{
		Width:  10,
		Height: 5,
	}
	fmt.Println(r)
	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp)
	fmt.Println(rp.area())
	fmt.Println(rp.perim())

	r.setWithValue("Width", 20)
	r.setWithPtr("Height", 100)
	fmt.Println(r)
	fmt.Println(r.area())
	fmt.Println(r.perim())

	return nil
}

func (m *MethodExpl) Init() {
	if err := goexpl.RegisterGoExample(m.GetGoExampleName(), &MethodExpl{}); err != nil {
		panic(err)
	}
}

func (m *MethodExpl) GetGoExampleName() string {
	return common.MethodExpl
}
