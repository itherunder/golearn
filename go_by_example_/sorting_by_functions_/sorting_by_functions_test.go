package sorting_by_functions_

import (
	"sort"
	"testing"

	"github.com/yezihack/colorlog"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func TestSortingByFunctions(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	colorlog.Info("Sorted: %v", fruits)
}
