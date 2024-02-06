package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// ChannelDirectionExpl
// channel can be declare with direction, you can specify a channel tobe send-only or receive-only
type ChannelDirectionExpl struct {
}

func (cde *ChannelDirectionExpl) RunExample(inputParams *goexpl.InputParams) error {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	return nil
}

func (cde *ChannelDirectionExpl) Init() {
	if err := goexpl.RegisterGoExample(cde.GetGoExampleName(), &ChannelDirectionExpl{}); err != nil {
		panic(err)
	}
}

func (cde *ChannelDirectionExpl) GetGoExampleName() string {
	return common.ChannelDirection
}
