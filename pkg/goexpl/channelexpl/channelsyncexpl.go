package channelexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
)

// ChannelSyncExpl
// channel can be used to synchronize execution across goroutine
// in this example, we use a blocking receive to wait a goroutine to finish
// if there are multiple goroutines to wait, please refer to WaitGroup
type ChannelSyncExpl struct {
}

func (cse *ChannelSyncExpl) RunExample(inputParams *goexpl.InputParams) error {
	done := make(chan bool, 1)
	// start a goroutine and give it a Channel to notify on
	go worker(done)

	// the code will be blocked here until receive a value from chan done
	<-done
	return nil
}

func (cse *ChannelSyncExpl) Init() {
	if err := goexpl.RegisterGoExample(cse.GetGoExampleName(), &ChannelSyncExpl{}); err != nil {
		panic(err)
	}
}

func (cse *ChannelSyncExpl) GetGoExampleName() string {
	return common.ChannelSync
}
