package goroutineexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// goroutine is a lightweight thread of execution
type GoRoutineExpl struct {
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func (gr *GoRoutineExpl) RunExample(inputParams *goexpl.InputParams) error {

	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going anonymous")

	time.Sleep(time.Second)
	fmt.Println("done")

	return nil
}

func (gr *GoRoutineExpl) Init() {
	if err := goexpl.RegisterGoExample(gr.GetGoExampleName(), &GoRoutineExpl{}); err != nil {
		panic(err)
	}
}

func (gr *GoRoutineExpl) GetGoExampleName() string {
	return common.GoRoutineExpl
}
