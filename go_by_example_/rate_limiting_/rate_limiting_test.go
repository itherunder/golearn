package rate_limiting_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestRateLimiting(t *testing.T) {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// time.Tick 和time.NewTicker 的区别：
	// http://www.classical.pub/archives/37/
	// 就是没啥区别...
	limiter := time.Tick(time.Millisecond * 500)

	for req := range requests {
		<-limiter
		colorlog.Info("request %d", req)
	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		colorlog.Info("request %v", req)
	}
}

func TestTickAndNewTicker(t *testing.T) {
	tick := time.Tick(time.Millisecond * 500)
	for {
		select {
		case t := <-tick:
			colorlog.Info("Tick at %v", t)
		}
	}

	ticker := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case t := <-ticker.C:
			colorlog.Info("Tick at %v", t)
		}
	}
}
