package structexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"reflect"
)

// golang 支持结构体和接口的嵌套
type Embedding struct {
}

// base struct
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v\ntype: %v", b.num, reflect.TypeOf(&b).Elem().Name())
}

// sub struct container embed a base.
// looks like a field without a name
type container struct {
	base
	str string
}

func (c container) getStr() string {
	return c.str
}

func (e *Embedding) RunExample(inputParams *goexpl.InputParams) error {

	// 新建一个嵌套了 base 结构体的变量时，需要同时为嵌套的 base 中的字段赋值。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}
	// co 结构体可以直接访问 base 的字段
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	// 也可以通过 base 访问 num 字段
	fmt.Println("also num:", co.base.num)

	// co 也可以调用 base 的方法
	fmt.Println(co.describe())

	// getstr
	fmt.Println(co.getStr())

	// 定一个接口，将方法嵌入该结构体，可以将该方法传递给其他结构体
	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	return nil
}

func (e *Embedding) Init() {
	if err := goexpl.RegisterGoExample(e.GetGoExampleName(), &Embedding{}); err != nil {
		panic(err)
	}
}

func (e *Embedding) GetGoExampleName() string {
	return common.Embed
}
