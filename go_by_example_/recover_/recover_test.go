package recover_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func MyPanic() {
	panic("a problem")
}

func TestRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			colorlog.Info("recovered: %v", r)
		}
	}()

	MyPanic()
	colorlog.Info("after mypanic")
}
