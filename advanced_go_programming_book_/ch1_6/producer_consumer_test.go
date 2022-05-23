package ch1_6

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		out <- factor * i
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func TestProducerConsumer(t *testing.T) {
	ch := make(chan int, 64) // buffered channel
	go Producer(2, ch)
	go Producer(5, ch)
	go Consumer(ch)

	// time.Sleep(5 * time.Second)

	// ctrl + c quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
