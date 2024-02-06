package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// ChannelBufferExpl
// here is a example for a channel, which has a buffer with certain capacity
type ChannelBufferExpl struct {
}

func (cbe *ChannelBufferExpl) RunExample(inputParams *goexpl.InputParams) error {
	// here we declare a channer with a buffer, whose capacity is 2
	messages := make(chan string, 2)

	messages <- "bufferd"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	close(messages)
	return nil
}

func (cbe *ChannelBufferExpl) Init() {
	if err := goexpl.RegisterGoExample(cbe.GetGoExampleName(), &ChannelBufferExpl{}); err != nil {
		panic(err)
	}
}

func (cbe *ChannelBufferExpl) GetGoExampleName() string {
	return common.ChannelBuffer
}
