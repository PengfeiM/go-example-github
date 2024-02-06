// Package structexpl
/*
	Go 的结构体是类型化的字段集合。它们对于将数据分组在一起形成记录非常有用。
*/

package structexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// 结构体 struct 的 golang example
type StructExpl struct {
}

// person 结构体，用于结构体 example
// 包含字段：
//	* name
//	* age
type person struct {
	name string
	age  uint
}

// newPerson
// 构造一个新的 person 结构体，并返回指针
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 18
	return &p
}

func (s *StructExpl) RunExample(inputParams *goexpl.InputParams) error {

	// the simple way to create a struct of person
	fmt.Println(person{"Bob", 20})

	// 也可以在创建 person 的时候指定参数名
	fmt.Println(person{name: "Alice", age: 30})

	// 未传入值的字段将置为 0 值
	fmt.Println(person{name: "Fred"})

	// 加上 & 将会返回一个指针
	fmt.Println(&person{name: "Ann", age: 40})

	// 通常，我们会把结构体的初始化放在一个 newType() 的函数中
	fmt.Println(newPerson("Jon"))

	ps := person{name: "Sean", age: 50}
	fmt.Println(ps.name) // use dot(.) to access a field of struct

	// use pointer can do the some thing, golang will dereference the pointer automatically
	psp := &ps
	fmt.Println(psp.age)

	psp.age = 81
	fmt.Println(psp.age)

	// 匿名结构体也是支持的
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	return nil
}

func (s *StructExpl) Init() {
	if err := goexpl.RegisterGoExample(s.GetGoExampleName(), &StructExpl{}); err != nil {
		panic(err)
	}
}

func (s *StructExpl) GetGoExampleName() string {
	return common.StructExpl
}
