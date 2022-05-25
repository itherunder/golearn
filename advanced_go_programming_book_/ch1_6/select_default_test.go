package ch1_6

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func receive(in <-chan string) {
	for {
		select {
		case v := <-in:
			fmt.Println(v)
		default: // will always run when no data
			// no data
			fmt.Println("default")
		}
	}
}

func TestSelectDefault(t *testing.T) {
	in := make(chan string)
	go receive(in)
	for i := 0; i < 10; i++ {
		in <- strconv.Itoa(i)
	}
	time.Sleep(2 * time.Second)
}
