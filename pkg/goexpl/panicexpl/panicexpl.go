package panicexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"os"
)

// PanicExpl
// Go's panic function is used to cause a run-time panic.
type PanicExpl struct{}

func (pe *PanicExpl) RunExample(inputParams *goexpl.InputParams) error {
	//panic("a problem")

	_, err := os.Create("./files/file")
	if err != nil {
		panic(err)
	}

	return nil
}

func (pe *PanicExpl) Init() {
	if err := goexpl.RegisterGoExample(pe.GetGoExampleName(), &PanicExpl{}); err != nil {
		panic(err)
	}
}

func (pe *PanicExpl) GetGoExampleName() string {
	return common.PanicExpl
}
