package slice

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestSlice(t *testing.T) {
	var str string = "eego世界"
	colorlog.Debug("str length: %v", len(str)) // str length: 10
}
