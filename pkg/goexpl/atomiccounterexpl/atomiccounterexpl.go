package atomiccounterexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"sync"
	"sync/atomic"
)

// AtomicCounterExpl
// apart from channel, there is another way to manage state across goroutines, which is using sync/atomic package
// for atomic counters accessed by multiple goroutines
type AtomicCounterExpl struct {
}

func (ace *AtomicCounterExpl) RunExample(inputParams *goexpl.InputParams) error {

	// declare a atomic value
	var ops uint64
	var cal uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		// start 50 goroutines to add 1 to ops repeatedly
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				// atomic add for value stored in ops
				atomic.AddUint64(&ops, 1)
				// a non-atomic add
				cal++
			}
			wg.Done()
		}()
	}

	// an atomic read
	fmt.Println(atomic.LoadUint64(&ops))
	wg.Wait()

	// a usual read
	fmt.Println(ops) // will be 50000
	fmt.Println(cal) // not sure the value of cal, since its addition is not atomic

	return nil
}

func (ace *AtomicCounterExpl) Init() {
	if err := goexpl.RegisterGoExample(ace.GetGoExampleName(), &AtomicCounterExpl{}); err != nil {
		panic(err)
	}
}

func (ace *AtomicCounterExpl) GetGoExampleName() string {
	return common.AtomicCounterExpl
}
