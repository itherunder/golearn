package sync_

import (
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/yezihack/colorlog"
)

func getGoRoutine() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// get id
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idField)
	return id
}

func TestMutex(t *testing.T) {
	var lock sync.Mutex
	go func() {
		lock.Lock()
		defer lock.Unlock()

		colorlog.Info("goroutine %v get lock", getGoRoutine())
		time.Sleep(time.Second)
		colorlog.Info("goroutine %v get unlock", getGoRoutine())
	}()

	time.Sleep(time.Second / 10)
	go func() {
		lock.Lock()
		defer lock.Unlock()

		colorlog.Info("goroutine %v get lock", getGoRoutine())
		time.Sleep(time.Second)
		colorlog.Info("goroutine %v get unlock", getGoRoutine())
	}()
	time.Sleep(time.Second * 4)
}

// RWMutext allow multiple read and single write
func TestRWMutex(t *testing.T) {
	var rwLock sync.RWMutex
	// get rlock
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.RLock()
			defer rwLock.RUnlock()
			colorlog.Info("read func %v get rlock at %v", i, time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second / 10)

	// get wlock
	for i := 0; i < 5; i++ {
		go func(i int) {
			rwLock.Lock()
			defer rwLock.Unlock()
			colorlog.Info("write func %v get wlock at %v", i, time.Now().String())
			time.Sleep(time.Second)
		}(i)
	}

	time.Sleep(time.Second * 10)
}

func TestMap(t *testing.T) {
	var mp sync.Map
	var wg sync.WaitGroup

	addNumber := func(begin int) {
		for i := begin; i < begin+3; i++ {
			mp.Store(i, i)
		}
		wg.Done()
	}

	routineSize := 5
	wg.Add(routineSize)
	for i := 0; i < routineSize; i++ {
		go addNumber(i * 10)
	}

	wg.Wait()
	var sz int
	// random list
	mp.Range(func(k, v interface{}) bool {
		sz++
		colorlog.Info("key: %+v, value: %+v", k, v)
		return true
	})
	colorlog.Info("size: %v", sz)
	v, ok := mp.Load(0)
	if ok {
		colorlog.Info("key 0 has value: %+v", v)
	}
}
