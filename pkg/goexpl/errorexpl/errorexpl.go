package errorexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// ErrorExpl
/*
golang 的 Error 示例
在 golang 中，通常使用独立的返回值来传递错误
这与 Java 和 Ruby 等语言中使用的异常以及 C 中有时使用的重载单个结果/错误值形成对比。Go 的方法可以轻松查看哪些函数返
回错误，并使用与任何其他函数相同的语言结构来处理它们，无错误任务
*/
type ErrorExpl struct {
}

func (e *ErrorExpl) RunExample(inputParams *goexpl.InputParams) error {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, err := f2(42)
	// how to access a self-defined error type data
	/*

	 */
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

	return nil
}

func (e *ErrorExpl) Init() {
	err := goexpl.RegisterGoExample(e.GetGoExampleName(), &ErrorExpl{})
	if err != nil {
		panic(err)
	}
}

func (e *ErrorExpl) GetGoExampleName() string {
	return common.ErrorExpl
}
