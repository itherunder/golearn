package tickers_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestTickers(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				colorlog.Info("Tick at %v", t)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	done <- true
	colorlog.Info("Ticker stopped")
}
