// Package mapexpl
/*
This is package for map exmples
*/

package mapexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type MapExpl struct {
}

func (me *MapExpl) RunExample(inputParams *goexpl.InputParams) error {
	/*
		这里我们进行 map 这个类型的 Example 测试
	*/

	// declare a map, which key is string, value is int
	var m map[string]int
	// map has no attribute cap, just len
	fmt.Println("len: ", len(m))
	// if you run the line bellow, the program will panic, because you haven't allocate a space to map
	//m["k1"] = 1

	// initialize a map
	m = make(map[string]int)
	// as you see, the length of m is still 0.
	fmt.Println("len: ", len(m))

	// here we add 2 k-v pairs into m
	m["k0"] = 0
	m["k1"] = 1
	m["k2"] = 12
	fmt.Println("map: ", m)

	// you can get value by key
	v1 := m["k1"]
	fmt.Println(v1)

	// In fact, it will return 2 vars when you try to get value by key from a map
	// as you can see the line bellow, the 2 return values are:
	//	- the value of the key, zero-value for non-exist key
	//	- a flag to show if your operation is successful
	// it would be useful to help you deal with access a key which is not in the map
	//v3, ok := m["k2"]
	v3, ok := m["k3"]
	if !ok {
		fmt.Println("the key is not exist, return 0-value")
		fmt.Println("v3: ", v3, "    ok: ", ok)
	} else {
		fmt.Println("get the value of the key successfully")
		fmt.Println("v3: ", v3, "    ok: ", ok)
	}

	// 内置的 delete(map, key) 函数可以用于从 map 中删除 k-v 对
	delete(m, "k1")
	fmt.Println("map: ", m)

	// 内置的 clear() 函数可以用于清空 map
	//clear(m)

	// 你也可以是用其他方式声明并初始化一个 map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map n: ", n)

	// 如果两个 map 要比较呢？
	n1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("is addr the same? ", &n == &n1)
	fmt.Printf("n addr:  %p\n", &n)
	fmt.Printf("n1 addr: %p\n", &n1)
	// 我们不能使用 == 来判断两个 map
	//fmt.Println("can we use ==? ", n == n1)
	// Equal 方法可以用，但是我这里 golang 1.18 暂时没有这个库
	//fmt.Println("maps.Equal(n, n1)? ", maps.Equal(n, n1))

	return nil
}

func (me *MapExpl) Init() {
	if err := goexpl.RegisterGoExample(me.GetGoExampleName(), &MapExpl{}); err != nil {
		panic(err)
	}
}

func (me *MapExpl) GetGoExampleName() string {
	return common.MapExpl
}
