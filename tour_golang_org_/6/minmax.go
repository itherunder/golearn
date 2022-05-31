package test6

import "fmt"

func test7() {
	var min, max int
	min, max = MinMax(78, 67)
	fmt.Printf("min is %d, max is %d\n", min, max)
}

func MinMax(a, b int) (min, max int) {
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	return
}
