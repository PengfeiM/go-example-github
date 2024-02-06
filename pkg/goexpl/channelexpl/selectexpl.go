package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// SelectExpl
// golang's select lets you wait on multiple channel operations.
// combine goroutine and channel with select is a powerfule feature in Go
type SelectExpl struct {
}

func (se *SelectExpl) RunExample(inputParamse *goexpl.InputParams) error {

	fmt.Println("use two Channel and goroutine to demonstrate this")
	// we create 2 channels here
	c1 := make(chan string)
	c2 := make(chan string)

	// send values into 2 channels after a certain time
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received:", msg1)
		case msg2 := <-c2:
			fmt.Println("received:", msg2)
		}
	}

	close(c1)
	close(c2)

	// here we use a death loop to demonstrate this feature
	fmt.Println("use select to echo input")
	msgs := make(chan string, 1)
	outputted := make(chan bool, 1)
	done := make(chan bool, 1)

	// start goroutine to monitor input
	go echor(msgs, outputted, done)

	// use select to receive values from msgs and done
OuterLoop:
	for {
		select {
		case output := <-msgs:
			fmt.Println("your input is:", output)
			outputted <- true
		case <-done:
			break OuterLoop
		}
	}

	close(msgs)
	close(outputted)
	close(done)

	return nil
}

func (se *SelectExpl) Init() {
	if err := goexpl.RegisterGoExample(se.GetGoExampleName(), &SelectExpl{}); err != nil {
		panic(err)
	}
}

func (se *SelectExpl) GetGoExampleName() string {
	return common.SelectExpl
}
