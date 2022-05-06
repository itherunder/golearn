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
	colorlog.Info("after mypanic") // 这行代码不会执行，因为 mayPanic 函数会调用 panic。 main 程序的执行在 panic 点停止，并在继续处理完 defer 后结束。
}
