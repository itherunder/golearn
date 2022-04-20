package list_

import (
	"container/list"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestList(t *testing.T) {
	tmpList := list.New()

	for i := 1; i <= 10; i++ {
		tmpList.PushBack(i)
	}

	first := tmpList.PushFront(0)
	tmpList.Remove(first)

	for l := tmpList.Front(); l != nil; l = l.Next() {
		colorlog.Info("l.Value: %v", l.Value)
	}
}
