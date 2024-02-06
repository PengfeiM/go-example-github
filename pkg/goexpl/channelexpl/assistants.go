package channelexpl

import (
	"fmt"
	"time"
)

// worker
// a func serve channelsyncexpl
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// when this worker finish it's task, it push a value into channel done
	done <- true
}

// ping
// a func, in which the chan is only support for push
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong
// The pong function accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func echor(msgs chan string, outputted chan bool, done chan bool) {

	fmt.Println("This is a long term input checker")

	var input string
	for {
		fmt.Println("Input something:")
		fmt.Scanln(&input)
		switch input {
		case "q", "quit":
			done <- true
			return
		default:
			msgs <- input
		}

		<-outputted
	}
}
