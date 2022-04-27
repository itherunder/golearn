package select_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestSelect(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()
	
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- ch1:
			colorlog.Info("received %v", msg1)
		case msg2 := <- ch2:
			colorlog.Info("received %v", msg2)
		}
	}
}