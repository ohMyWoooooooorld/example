package _func

import (
	"fmt"
	"time"
)

func RateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	// 从 limiter 接收一个值，然后每 200 毫秒接收一个值
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	chanLimiter := make(chan time.Time, 3)

	for range 3 {
		chanLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			chanLimiter <- t
		}
	}()

	chanRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		chanRequests <- i
	}
	close(chanRequests)
	for req := range chanRequests {
		<-chanLimiter
		fmt.Println("request", req, time.Now())
	}

}
