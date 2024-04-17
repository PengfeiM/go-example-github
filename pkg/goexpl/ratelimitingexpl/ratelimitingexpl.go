package ratelimitingexpl

import (
	"fmt"
	"go-example/pkg/common"
	"go-example/pkg/goexpl"
	"time"
)

// RateLimitingExpl
// we can implement rate limiting elegantly in go by goroutines, channels, tickers
type RateLimitingExpl struct {
}

func (rle *RateLimitingExpl) RunExample(inputParams *goexpl.InputParams) error {
	// 1st, implement a basic rate limiter
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter chan will receive value by a interval of 200ms
	// a leaky bucket method
	limiter := time.NewTicker(1000 * time.Millisecond)

	// process req when a msg popped out from leaky bucket limiter
	fmt.Println("use time.ticker and chan to implement a leaky bucket rate limiter")
	for req := range requests {
		<-limiter.C
		fmt.Println("request:", req, time.Now())
	}

	// next, we will implement a token bucket, which support a short burst
	// a chan
	burstLimiter := make(chan time.Time, 3)

	// initialize the token bucket
	for i := 0; i < 3; i++ {
		burstLimiter <- time.Now()
	}

	// push tokens to bucket every 1000ms
	go func() {
		for t := range time.NewTicker(1000 * time.Millisecond).C {
			burstLimiter <- t
		}
	}()

	burstRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstRequests <- i
	}
	close(burstRequests)
	fmt.Println("a demonstration of token bucket")
	for req := range burstRequests {
		// do req when you can get a token from bucket
		<-burstLimiter
		fmt.Println("request:", req, time.Now())
	}

	return nil
}

func (rle *RateLimitingExpl) Init() {
	if err := goexpl.RegisterGoExample(rle.GetGoExampleName(), &RateLimitingExpl{}); err != nil {
		panic(err)
	}
}

func (rle *RateLimitingExpl) GetGoExampleName() string {
	return common.RateLimitingExpl
}
