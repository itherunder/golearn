package sorting_

import (
	"sort"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestSorting(t *testing.T) {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	colorlog.Info("Strings sorted: %v", strs)

	ints := []int{5, 3, 3, 2, 1}
	sort.Ints(ints)
	colorlog.Info("Ints sorted: %v", ints)

	s := sort.IntsAreSorted(ints)
	colorlog.Info("Sorted: %v", s)
}
