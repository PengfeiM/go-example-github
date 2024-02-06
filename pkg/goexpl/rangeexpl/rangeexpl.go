package rangeexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type RangeExpl struct {
}

func (r *RangeExpl) RunExample(inputParams *goexpl.InputParams) error {

	// init a slice of int
	nums := []int{2, 3, 4}
	sum := 0
	// for range 每次返回一个元素和它的index。
	// index, element 的格式。
	// 如下所示，我们丢弃了 index，将 element 赋值给 num
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 遍历 slice，获取特定元素的 index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range 也可以遍历 map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// 或者只遍历 keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// 也可以用于遍历 string，这一块具体参考 string 和 rune 部分
	for i, c := range "go" {
		fmt.Println(i, c)
	}
	return nil
}

func (r *RangeExpl) Init() {
	if err := goexpl.RegisterGoExample(r.GetGoExampleName(), &RangeExpl{}); err != nil {
		panic(err)
	}
}

func (r *RangeExpl) GetGoExampleName() string {
	return common.RangeExpl
}
