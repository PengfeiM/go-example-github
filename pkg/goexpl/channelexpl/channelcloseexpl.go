package channelexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// ChannelCloseExpl
// a example to test close channel
type ChannelCloseExpl struct {
}

func (c ChannelCloseExpl) RunExample(inputParams *goexpl.InputParams) error {
	// make 2 chan, one for jobs, one for done
	jobs := make(chan int, 5)
	done := make(chan bool)

	// start a goroutine receive job from jobs chan
	// send true to done chan when job is finished
	go func() {
		for {
			j, more := <-jobs
			// more will be false when chan jobs is closed and there is no more values in jobs chan
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}

		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
		// use sleep to control sending's interval
		time.Sleep(time.Second)
	}
	fmt.Println("sent all jobs")
	fmt.Println("closing jobs")
	close(jobs) // even chan is closed, the goroutine above will continue until chan is empty
	fmt.Println("jobs closed")
	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)

	return nil
}

func (c ChannelCloseExpl) Init() {
	if err := goexpl.RegisterGoExample(c.GetGoExampleName(), &ChannelCloseExpl{}); err != nil {
		panic(err)
	}
}

func (c ChannelCloseExpl) GetGoExampleName() string {
	return common.ChannelCloseExpl
}
