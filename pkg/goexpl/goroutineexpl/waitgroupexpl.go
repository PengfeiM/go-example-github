package goroutineexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"sync"
	"time"
)

type WaitGroupExpl struct {
}

func (wge *WaitGroupExpl) RunExample(inputParams *goexpl.InputParams) error {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	return nil
}

func (wge *WaitGroupExpl) Init() {
	if err := goexpl.RegisterGoExample(wge.GetGoExampleName(), &WaitGroupExpl{}); err != nil {
		panic(err)
	}
}

func (wge *WaitGroupExpl) GetGoExampleName() string {
	return common.WaitGroupExpl
}

func worker(id int) {
	fmt.Printf("worker %d start\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}
