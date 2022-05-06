package string_formatting_

import (
	"fmt"
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

type point struct {
	x, y int
}

var info = colorlog.Info

func TestStringFormatting(t *testing.T) {
	p := point{1, 2}
	info("struct1: %v\n", p)

	info("struct2: %+v\n", p)

	info("struct3: %#v\n", p)

	info("type: %T\n", p)

	info("bool: %t\n", true)

	info("int: %d\n", 123)

	info("bin: %b\n", 14)

	info("char: %c\n", 33)

	info("hex: %x\n", 456)

	info("float1: %f\n", 78.9)

	info("float2: %e\n", 123400000.0)
	info("float3: %E\n", 123400000.0)

	info("str1: %s\n", "\"string\"")

	info("str2: %q\n", "\"string\"")

	info("str3: %x\n", "hex this")

	info("pointer: %p\n", &p)

	info("width1: |%6d|%6d|\n", 12, 345)

	info("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	info("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	info("width4: |%6s|%6s|\n", "foo", "b")

	info("width5: |%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
	fmt.Fprintf(os.Stdout, "io: an %s\n", "stdout")
}
