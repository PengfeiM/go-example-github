package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

type ChannelRangeExpl struct {
}

func (cre *ChannelRangeExpl) RunExample(input *goexpl.InputParams) error {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// this loop will end when the channel is closed and the channel is empty
	for elem := range queue {
		fmt.Println(elem)
	}

	return nil
}

func (cre *ChannelRangeExpl) Init() {
	if err := goexpl.RegisterGoExample(cre.GetGoExampleName(), &ChannelRangeExpl{}); err != nil {
		panic(err)
	}
}

func (cre *ChannelRangeExpl) GetGoExampleName() string {
	return common.ChannelRangeExpl
}
