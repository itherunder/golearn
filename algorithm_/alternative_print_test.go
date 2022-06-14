package algorithm_

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/yezihack/colorlog"
)

func curGoRoutine() int {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine ")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return id
}

func alternativePrint(x, n int) {
	chs := make([]chan bool, x)
	var wg sync.WaitGroup
	wg.Add(x)
	defer wg.Wait()
	for i := 0; i < x; i++ {
		chs[i] = make(chan bool)
	}
	for i := 0; i < x; i++ {
		go func(i int) {
			defer wg.Done()
			for j := i + 1; j <= n; j += x {
				<-chs[i]
				colorlog.Info("%d: %d", curGoRoutine(), j)
				if j < n {
					chs[(i+1)%x] <- true
				}
			}
		}(i)
	}
	chs[0] <- true
}

func TestAlternativePrint(t *testing.T) {
	alternativePrint(35, 1131)
}
