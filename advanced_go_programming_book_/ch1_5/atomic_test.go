package ch1_5

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/yezihack/colorlog"
)

var total struct {
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func TestAtomic(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)

	go worker(&wg)
	go worker(&wg)

	wg.Wait()
	colorlog.Info("total: %d", total.value)
}

var total_ uint64

func worker1(wg *sync.WaitGroup) {
	defer wg.Done()

	var i uint64
	for ; i <= 100; i++ {
		atomic.AddUint64(&total_, i)
	}
}

func TestAtomic1(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker1(&wg)
	go worker1(&wg)

	wg.Wait()

	fmt.Println("total:", total_)
}
