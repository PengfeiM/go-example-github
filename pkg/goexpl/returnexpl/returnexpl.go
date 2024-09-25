package returnexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type ReturnExpl struct {
}

func (r *ReturnExpl) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println("a test for return func order\n\tcode: return func2(), func1()")
	fmt.Println(returnFunc())
	fmt.Println("as you see, when the func returns, it will execte func1 and func2 in order")

	fmt.Println("this is a test for \n\treturn out, func(){/*modify out*/}()")
	fmt.Println(modifyReturn())
	fmt.Println("as you see, when the func returns out, func(), it execute func and return out and func's return value")

	return nil
}

func (r *ReturnExpl) Init() {
	if err := goexpl.RegisterGoExample(r.GetGoExampleName(), &ReturnExpl{}); err != nil {
		panic(err)
	}
}

func (r *ReturnExpl) GetGoExampleName() string {
	return common.Return
}
