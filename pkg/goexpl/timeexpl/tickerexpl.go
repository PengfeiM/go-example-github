package timeexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// TickerExpl 定时器
// using timer, you can do something once in the future
// using ticker, you can do something repeatedly by a specified interval
type TickerExpl struct{}

func (te *TickerExpl) RunExample(input *goexpl.InputParams) error {

	ticker := time.NewTicker(500 * time.Millisecond)

	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	println("Ticker stopped")
	return nil
}

func (te *TickerExpl) Init() {
	if err := goexpl.RegisterGoExample(te.GetGoExampleName(), &TickerExpl{}); err != nil {
		panic(err)
	}
}

func (te *TickerExpl) GetGoExampleName() string {
	return common.TickerExpl
}
