package ch1_6

import (
	"fmt"
	"testing"
)

func TestSelectRandom(t *testing.T) {
	ch := make(chan int)
	size := 10
	go func() {
		for i := 0; i < size; i++ {
			select { // select will random choose one case
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	// for v := range ch { // doesn't work when for loop has the limit
	// 	fmt.Print(v)
	// }
	for i := 0; i < size; i++ {
		fmt.Print(<-ch)
	}
	fmt.Println()
}
