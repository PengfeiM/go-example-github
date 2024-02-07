package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// TimeOutExpl
// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution
// time.
// We Implement timeouts in golang with channel and select elegantly
type TimeOutExpl struct {
}

func (toe *TimeOutExpl) RunExample(inputParams *goexpl.InputParams) error {

	// make a channel and start a goroutine to send value to it
	c1 := make(chan string, 1)
	// here we start a goroutine to send a value into c1 after 2s.
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result: 1"
	}()

	// in this select, we implement a timeout.
	// res := <- c1 waits for a value send into c1
	// <-time.After awaits a value to be dent after timeout of 1s
	// considering you can only get a value for res after 2 sec
	// in this case, the program will go into time.After timeout case
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// if we use  a timeout of 3s, what will happen?
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result: 2"
	}()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	case res := <-c2:
		fmt.Println(res)
	}
	close(c1)
	close(c2)

	return nil
}

func (toe *TimeOutExpl) Init() {
	if err := goexpl.RegisterGoExample(toe.GetGoExampleName(), &TimeOutExpl{}); err != nil {
		panic(err)
	}
}

func (toe *TimeOutExpl) GetGoExampleName() string {
	return common.TimeOut
}
