package bytes_

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("Hello, ")
	buf.WriteByte('G')
	buf.WriteString("o")
	buf.Write([]byte{'l', 'o', 'n', 'g'})
	fmt.Printf("buf: %v\n", buf.String())
	fmt.Printf("ints: %v\n", intsToString([]int{1, 2, 3}))
}

func intsToString(vals []int) string {
	var buf bytes.Buffer
	for i, v := range vals {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(fmt.Sprintf("%d", v))
	}
	return buf.String()
}
