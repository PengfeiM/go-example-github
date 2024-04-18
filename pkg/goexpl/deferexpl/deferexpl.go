package deferexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type DeferExpl struct {
}

func (de *DeferExpl) RunExample(inputParams *goexpl.InputParams) error {

	f := createFile("./files/defer.txt")
	defer closeFile(f)
	writeFile(f)

	return nil
}

func (de *DeferExpl) Init() {
	if err := goexpl.RegisterGoExample(de.GetGoExampleName(), &DeferExpl{}); err != nil {
		panic(err)
	}
}

func (de *DeferExpl) GetGoExampleName() string {
	return common.DeferExpl
}
