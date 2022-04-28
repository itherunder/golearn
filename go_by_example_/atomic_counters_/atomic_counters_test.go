package atomic_counters_

import (
	"sync"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestAtomicCounters(t *testing.T) {
	var op uint64 = 0

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				// atomic.AddUint64(&op, 1) // 50000
				op += 1 // 49334
			}
			wg.Done()
		}()
	}
	wg.Wait()

	colorlog.Info("ops: %d", op)
}
