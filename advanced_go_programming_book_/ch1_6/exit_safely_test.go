package ch1_6

import (
	"fmt"
	"testing"
	"time"
)

func worker(ch chan bool) {
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
			// exit
		}
	}
}

func TestExitSafely(t *testing.T) {
	ch := make(chan bool)
	defer close(ch) // close channel will broadcast a nil value and an optional fail flag

	for i := 0; i < 10; i++ {
		go worker(ch)
	}

	time.Sleep(time.Second)
}
