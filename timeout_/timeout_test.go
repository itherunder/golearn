package timeout

import (
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

type DelayResponseInterface interface {
	Handler(http.ResponseWriter, *http.Request)
	DelayHandler(string) http.ResponseWriter
	Done(string) http.ResponseWriter
}

type DelayResponse struct {
	mu       sync.Mutex
	timeout  time.Duration
	requests map[string]chan *http.Request
	times    map[string]uint
}

func NewDelayResponse(timeout time.Duration) *DelayResponse {
	return &DelayResponse{
		timeout:  timeout,
		times:    make(map[string]uint),
		requests: make(map[string]chan *http.Request),
	}
}

func (dr *DelayResponse) DelayHandler(remote string, w *http.ResponseWriter) {
	for {
		select {
		case <-time.After(dr.timeout):
			dr.Done(remote, w)
			return
		case <-dr.requests[remote]:
			colorlog.Info("new continued request arrived: %v", dr.times[remote])
		}
	}
}

func (dr *DelayResponse) Handler(w http.ResponseWriter, r *http.Request) {
	dr.mu.Lock()
	defer dr.mu.Unlock()
	remote := r.RemoteAddr
	if _, ok := dr.times[remote]; !ok {
		dr.times[remote] = 1
		dr.requests[remote] = make(chan *http.Request)
		go dr.DelayHandler(remote, &w)
	} else {
		dr.times[remote]++
	}
	dr.requests[remote] <- r
}

func (dr *DelayResponse) Done(remote string, w *http.ResponseWriter) {
	dr.mu.Lock()
	defer dr.mu.Unlock()
	colorlog.Info("total request times: %v", dr.times[remote])
	(*w).Write([]byte(fmt.Sprintf("total request times: %v", dr.times[remote])))
	delete(dr.times, remote)
	delete(dr.requests, remote)
}

func TestTimeout(t *testing.T) {
	stop := make(chan bool)
	dr := NewDelayResponse(3 * time.Second)
	http.HandleFunc("/req", dr.Handler)
	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		colorlog.Info("stop")
		stop <- true
	})
	http.ListenAndServe("localhost:1234", nil)
	<-stop
}
