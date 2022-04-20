package rev_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func TestRev(t *testing.T) {
	s := []byte("itherunder is liaozhou")
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	colorlog.Info("\nbefore: %s\nafter: %s\n", []byte("itherunder is liaozhou"), s)

	a := [...]int{0, 1, 2, 3, 4, 5}
	colorlog.Info("a is %v", a)
	reverse(a[:2]) // left move
	reverse(a[2:])
	reverse(a[:])
	colorlog.Info("a is %v", a)
}

func reverse(s []int) {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
