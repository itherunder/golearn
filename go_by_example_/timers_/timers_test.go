package timers_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestTimers(t *testing.T) {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	colorlog.Info("Timer 1 expired")

	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C
		colorlog.Info("Timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		colorlog.Info("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)
}
