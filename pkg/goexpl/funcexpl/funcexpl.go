package funcexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type FuncExpl struct {
}

// 定义一个加法函数，返回 a 和 b 的和
// 函数定义格式：
// func funcName(param1 type, param2 type) returnType {}
func plus(a int, b int) int {
	return a + b
}

// 另外一种格式：
// func funcName(param1, param2, param3 type) returnType {}
func plusPlus(a, b, c int) int {
	return a + b + c
}

// 多个返回值也是可以支持的
// func funcName() (returnType1, returnType2) {}
func vals() (int, int) {
	return 3, 7
}

// 不定参数的函数可以传入任意数量的参数
// 比如 fmt.Println()
// 下面定义了一个不定数量 int 类型参数的函数。
// nums 是一个 []int 类型变量。
// func Func(slice ...type){}
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func (f *FuncExpl) RunExample(inputParams *goexpl.InputParams) error {

	// call func plus
	res := plus(1, 2)
	fmt.Println("1+2 = ", res)

	// call func plusPlus
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3= ", res)

	// call multi return value func
	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	_, c := vals()
	fmt.Println("c:", c)

	// 不定参数函数
	sum(1, 2)
	sum(1, 2, 3)
	sum([]int{1, 2, 3, 4}...)

	// 也可以定义一个匿名函数
	res = func(a int, b int) int {
		return a + b
	}(1, 2) // 在这里直接调用了该函数并将返回值赋值给 res
	fmt.Println("匿名函数：")
	fmt.Println("1+2 = ", res)

	// 也可以先定义函数在调用
	addFunc := func(a int, b int) int {
		return a + b
	}
	//fmt.Println("匿名函数：", addFunc) // addFunc 的地址会被打印
	res = addFunc(1, 2)
	fmt.Println("匿名函数（先声明，再调用）：")
	fmt.Println("1+2 = ", res)

	return nil
}

func (f *FuncExpl) Init() {
	if err := goexpl.RegisterGoExample(f.GetGoExampleName(), &FuncExpl{}); err != nil {
		panic(err)
	}
}

func (f *FuncExpl) GetGoExampleName() string {
	return common.FuncExpl
}
