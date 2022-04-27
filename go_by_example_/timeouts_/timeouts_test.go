package timeout_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestTimeout(t *testing.T) {
	ch1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "msg1"
	}()
	
	select {
	case msg1 := <- ch1:
		colorlog.Info("received msg1 %v", msg1)
	case t := <- time.After(1 * time.Second):
		colorlog.Info("msg1 timeout %v", t)
	}

	ch2 := make(chan string, 2)
	
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "msg2"
	}()
	
	select {
	case msg1 := <- ch1:
		colorlog.Info("received msg2 %v", msg1)
	case t := <- time.After(3 * time.Second):
		colorlog.Info("msg2 timeout %v", t)
	}
}