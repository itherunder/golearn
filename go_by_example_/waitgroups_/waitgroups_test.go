package waitgroups_

import (
	"sync"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func worker(id int) {
	colorlog.Info("Worker %d starting", id)
	time.Sleep(time.Second)
	colorlog.Info("Worker %d finished", id)
}

func TestWaitgroups(t *testing.T) {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 避免在每个协程闭包中重复利用相同的 i 值
		// 更多细节可以参考 [the FAQ](https://go.dev/doc/faq#closures_and_goroutines)
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
