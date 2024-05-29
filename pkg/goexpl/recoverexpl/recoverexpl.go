package recoverexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type RecoverExpl struct{}

func mayPanic() {
	panic("a problem")
}

func (re *RecoverExpl) RunExample(inputParams *goexpl.InputParams) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recoverd. Error:\n", r)
		}
	}()

	mayPanic()

	// this part of code can not be executed
	// because it's writed after the panic
	// the program panic and enter defer to recover
	fmt.Println("After mayPanic()")
	return nil
}

func (re *RecoverExpl) Init() {
	if err := goexpl.RegisterGoExample(re.GetGoExampleName(), &RecoverExpl{}); err != nil {
		panic(err)
	}
}

func (re *RecoverExpl) GetGoExampleName() string {
	return common.RecoverExpl
}
