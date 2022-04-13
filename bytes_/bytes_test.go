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
	fmt.Printf("comma: %v\n", comma(5233451))
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

func comma(i int) string {
	s := fmt.Sprintf("%d", i)
	var buf bytes.Buffer
	for i, c := len(s)-1, 0; i >= 0; i, c = i-1, c+1 {
		if c > 0 && c%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	res := buf.Bytes()
	l := len(res)
	for i := 0; i < l/2; i++ {
		res[i], res[l-i-1] = res[l-i-1], res[i]
	}
	return string(res)
}
