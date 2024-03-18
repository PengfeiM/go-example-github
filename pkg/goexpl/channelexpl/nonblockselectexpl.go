package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type NonBlockSelectExpl struct {
}

func (n *NonBlockSelectExpl) RunExample(inputParams *goexpl.InputParams) error {
	// declare 2 channels
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

	close(messages)
	close(signals)

	return nil
}

func (n *NonBlockSelectExpl) Init() {
	if err := goexpl.RegisterGoExample(n.GetGoExampleName(), &NonBlockSelectExpl{}); err != nil {
		panic(err)
	}
}

func (n *NonBlockSelectExpl) GetGoExampleName() string {
	return common.NonBlockSelectExpl
}
