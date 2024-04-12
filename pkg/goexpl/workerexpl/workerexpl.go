package workerexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// WorkerExpl
// we can implement a worker pool by using goroutines and channels
type WorkerExpl struct{}

func (we *WorkerExpl) RunExample(inputParams *goexpl.InputParams) error {
	var numOfJobs, numOfWorkers int
	if inputParams.GetArgNum() == 0 {
		numOfJobs = 5
		numOfWorkers = 3
	} else if inputParams.GetArgNum() != 2 {
		panic("number of arguments must be 2 integer")
	}

	// make 2 channel
	// jobs channel will be used to send jobs to the worker
	// results channel will be used to receive the results from the worker
	jobs := make(chan int, numOfJobs)
	results := make(chan int, numOfJobs)

	for w := 1; w <= numOfWorkers; w++ {
		go worker(w, jobs, results)
	}

	// send 5 jobs to the jobs channel
	for j := 1; j <= numOfJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// receive the results from the results channel
	for a := 1; a <= numOfJobs; a++ {
		<-results
	}

	return nil
}

func (we *WorkerExpl) Init() {
	if err := goexpl.RegisterGoExample(we.GetGoExampleName(), &WorkerExpl{}); err != nil {
		panic(err)
	}
}

func (we *WorkerExpl) GetGoExampleName() string {
	return common.WorkerExpl
}

// define a worker
// which will receive work from the jobs channel, do it and send the result to the results channel
// by an interval of 1 second
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
