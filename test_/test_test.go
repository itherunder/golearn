package test_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestTest(t *testing.T) {
	var i int32 = -2147483648
	colorlog.Info("%x", i)
}
