package goroutineexpl

import (
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"math/rand"
	"sync/atomic"
	"time"
)

// StatefulGoroutineExpl
// in this example, we use built-in synchronization features of goroutines and channels to synchronize access to
// shared state across multiple goroutines
type StatefulGoroutineExpl struct {
}

func (sge *StatefulGoroutineExpl) RunExample(inputParams *goexpl.InputParams) error {
	// In this example, our state will be owned by one goroutine, which will guarantee that the date is never corrupted
	// with concurrent access.
	// to read or write that state, other goroutines will send messages to the owning goroutine and receive
	// corresponding replies. these readOp and writeOp structs encapsulate those requests and a way for the owning
	// goroutine to respond.

	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	writeOpsFinal := atomic.LoadUint64(&writeOps)

	println("readOps:", readOpsFinal)
	println("writeOps:", writeOpsFinal)

	return nil
}

func (sge *StatefulGoroutineExpl) Init() {
	if err := goexpl.RegisterGoExample(sge.GetGoExampleName(), &StatefulGoroutineExpl{}); err != nil {
		panic(err)
	}
}

func (sge *StatefulGoroutineExpl) GetGoExampleName() string {
	return common.StatefulGoroutineExpl
}
