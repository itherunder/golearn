package worker_pools_

import (
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		colorlog.Info("Worker %d starting job %d", id, j)
		time.Sleep(time.Second)
		colorlog.Info("Worker %d finished job %d", id, j)
		results <- j * 2
	}
}

func TestWorkerPools(t *testing.T) {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	defer close(jobs)
	defer close(results)

	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for i := 0; i < numJobs; i++ {
		<-results
	}
}
