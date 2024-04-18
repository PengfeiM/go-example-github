package slice

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"golang.org/x/exp/slices"
)

type SortByFuncExpl struct{}

func (sfe *SortByFuncExpl) RunExample(inputParams *goexpl.InputParams) error {
	fruits := []string{"apple", "banana", "pear", "orange", "pineapple", "kiwi", "peach", "strawberry", "blueberry"}

	lenCmp := func(a, b string) int {
		//if len(a) > len(b) {
		//	return 1
		//} else if len(a) < len(b) {
		//	return -1
		//} else {
		//	return 0
		//}
		return len(a) - len(b)
	}

	slices.SortFunc(fruits, lenCmp)
	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		Person{"Jax", 37},
		Person{"Bob", 31},
		Person{"Joe", 45},
		Person{"Dan", 27},
		Person{"Anne", 21},
	}

	slices.SortFunc(people,
		func(a, b Person) int {
			return a.age - b.age
		})
	fmt.Println(people)

	return nil
}

func (sfe *SortByFuncExpl) Init() {
	if err := goexpl.RegisterGoExample(sfe.GetGoExampleName(), &SortByFuncExpl{}); err != nil {
		panic(err)
	}
}

func (sfe *SortByFuncExpl) GetGoExampleName() string {
	return common.SortByFunc
}
