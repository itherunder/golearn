package ch1_5

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func TestConfigWithAtomicValue(t *testing.T) {
	var config atomic.Value
	config.Store(loadConfig())

	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
		}
	}()

	for i := 0; i < 10; i++ {
		go func() {
			for _, r := range requests() {
				time.Sleep(time.Second)
				c := config.Load().(string)
				// ...
				colorlog.Info("current config: %s", c)
				colorlog.Info("handle request: %s", r)
			}
		}()
	}
	time.Sleep(15 * time.Second)
}

func loadConfig() string {
	return fmt.Sprintf("current time: %d", time.Now().Unix())
}

func requests() []string {
	return []string{"a", "b", "c"}
}
