package channel_synchronization_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func worker(done chan bool) {
	colorlog.Info("working...")
	time.Sleep(2 * time.Second)
	colorlog.Info("done")
	done <- true
}

func TestChannelSync(t *testing.T) {
	done := make(chan bool, 1)
	go worker(done)

	<- done
}
