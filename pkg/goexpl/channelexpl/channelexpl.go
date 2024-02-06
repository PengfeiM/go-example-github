package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

/*
ChannelExpl
	channel 是 golang 中一种沟通并行的 goroutines 的一种 pipe/管道
	you can send values into channel from one goroutine,
	and receive these values into another goroutine
	FYI, a default channel has no buffer, which means that you can only push a values into it when the values in it is
	taken somewhere else. A channel with buffer will be discussed in ChannelBufferExpl
	In another way, it means no buffer channel can only receive values if there is a corresponding receive
*/
type ChannelExpl struct {
}

func (c *ChannelExpl) RunExample(inputParams *goexpl.InputParams) error {
	// how to declare a channel
	var msgChan chan string
	// before init, a channel is nil
	fmt.Println(msgChan)
	// init a channel
	msgChan = make(chan string)

	// start a goroutine to push a string to msgChan
	go func() {
		msgChan <- "ping"
	}()

	// receive string from msgChan in main goroutine
	msg := <-msgChan
	fmt.Println(msg)

	close(msgChan)
	return nil
}

func (c *ChannelExpl) Init() {
	if err := goexpl.RegisterGoExample(c.GetGoExampleName(), &ChannelExpl{}); err != nil {
		panic(err)
	}
}

func (c *ChannelExpl) GetGoExampleName() string {
	return common.Channel
}
