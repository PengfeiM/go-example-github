package timeexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// TimerExpl 定时器
// use timer to execute some code after some time
type TimerExpl struct{}

func (te *TimerExpl) RunExample(inputParams *goexpl.InputParams) error {

	// time1
	// timers are used to execute some code after some time, which you tell to the timer
	// it provides a channel to notify the timer when it expires
	fmt.Println(time.Now())
	timer1 := time.NewTimer(2 * time.Second) // a 2 sec timer

	// code will be blocked here until timer expires
	<-timer1.C
	fmt.Println("Timer1 fired")
	fmt.Println(time.Now())

	// the difference between timer and time.Sleep is that timer is cancelable,
	// so you can stop the timer before it expires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	fmt.Println(time.Now())
	stop2 := timer2.Stop()
	if stop2 {
		// timer2 is stopped before it fired
		fmt.Println(time.Now())
		fmt.Println("Timer 2 stopped")
	}

	return nil
}

func (te *TimerExpl) Init() {
	if err := goexpl.RegisterGoExample(te.GetGoExampleName(), &TimerExpl{}); err != nil {
		panic(err)
	}
}

func (te *TimerExpl) GetGoExampleName() string {
	return common.TimerExpl
}
