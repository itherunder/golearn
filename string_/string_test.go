package string_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestByteAndRune(t *testing.T) {
	str := "hello world"
	colorlog.Debug("str length: %v", len(str)) // str length: 11
	for i, v := range str {                    // type rune is int32
		colorlog.Debug("str[%v] = %c, type %T", i, v, v)
	}
}
