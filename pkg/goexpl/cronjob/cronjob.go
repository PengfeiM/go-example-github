package cronjob

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"

	"github.com/robfig/cron/v3"
)

type CronJob struct {
}

func (cj *CronJob) GetGoExampleName() string {
	return common.CronJob
}

var (
	spec   = "@every 1s" // spec, used to schedule job
	count1 = 0           // how many times a func1 run, it will panic when count == 1
	count2 = 0           // how many times a func1 run, it will panic when count == 1
)

func (cj *CronJob) RunExample(inputParams *goexpl.InputParams) error {
	fmt.Println("# Run example CronJob...")

	// make 2 cron job, one with recovery, one's not
	c1 := cron.New()                                                 // default cronjob
	c2 := cron.New(cron.WithChain(cron.Recover(cron.DefaultLogger))) //cronjob with recovery

	// add func
	if _, err := c1.AddFunc(spec, c1Worker); err != nil {
		return err
	}
	if _, err := c2.AddFunc(spec, c2Worker); err != nil {
		return err
	}

	// start job, choose one to run
	//c1.Start() //this one is going to panic and make the program exit
	c2.Start() // this doesn't

	// wait certain sec to watch cronjob
	time.Sleep(5 * time.Second)

	fmt.Println("# Example Over! ")
	return nil
}

func c1Worker() {
	count1++
	if count1 == 1 {
		panic("Ooops, this is a designed panic")
	}
	fmt.Printf("c1: Hi! for %d time(s)\n", count1)
}

func c2Worker() {
	count2++
	if count2 == 1 {
		panic("Ooops, this is a designed panic")
	}
	fmt.Printf("c2: Hello! for %d time(s)\n", count2)
}

func (cj *CronJob) Init() {
	if err := goexpl.RegisterGoExample(cj.GetGoExampleName(), &CronJob{}); err != nil {
		panic(err)
	}
}
