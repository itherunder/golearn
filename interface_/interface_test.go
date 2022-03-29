package interface_

import (
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestInterface(t *testing.T) {
	var e interface{}
	f, err := os.Open("./test.txt")
	if err != nil {
		colorlog.Error("open file error: %v", err)
	}
	e = f
	colorlog.Info("e type: %T", e)
	r, ok := e.(*os.File)
	if ok {
		colorlog.Info("r type: %T", r)
	}
}
