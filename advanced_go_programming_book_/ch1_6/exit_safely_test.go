package ch1_6

import (
	"fmt"
	"sync"
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

func worker1(wg *sync.WaitGroup, ch chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-ch:
			return
		}
	}
}

func TestExitSafely1(t *testing.T) {
	ch := make(chan bool)
	defer close(ch)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker1(&wg, ch)
	}

	time.Sleep(time.Second)
	wg.Wait()
}
