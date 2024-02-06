package helloworld

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type HelloWorld struct {
}

func (hw *HelloWorld) GetGoExampleName() string {
	return common.HelloWorld
}

func (hw *HelloWorld) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println("Hello, World!")
	return nil
}

func (hw *HelloWorld) Init() {
	err := goexpl.RegisterGoExample(hw.GetGoExampleName(), &HelloWorld{})
	if err != nil {
		panic(err)
	}
}
