package mutexes_

import (
	"sync"
	"testing"

	"github.com/yezihack/colorlog"
)

var globalMu sync.Mutex

// 由于map 不是线程安全的，所以在多个协程中访问同一个 map 时，报错
type Container struct {
	// mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// c.mu.Lock()
	// defer c.mu.Unlock()
	globalMu.Lock()
	defer globalMu.Unlock()
	c.counters[name]++
}

func TestMutexes(t *testing.T) {
	c := &Container{counters: map[string]int{"a": 0, "b": 0}}
	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	colorlog.Info("c.counters: %+v", c.counters)
}
