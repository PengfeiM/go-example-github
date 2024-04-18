package mutexexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"sync"
)

// MutexExpl
// a example to show how to use mutex to access date safely across multiple goroutines, apart from atomic counter, which
// is a good way in simple situations
type MutexExpl struct{}

func (mte *MutexExpl) RunExample(inputParams *goexpl.InputParams) error {
	c := MutexContainer{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()

	fmt.Println(c.counters)

	return nil
}

func (mte *MutexExpl) Init() {
	if err := goexpl.RegisterGoExample(mte.GetGoExampleName(), &MutexExpl{}); err != nil {
		panic(err)
	}
}

func (mte *MutexExpl) GetGoExampleName() string {
	return common.MutexExpl
}
