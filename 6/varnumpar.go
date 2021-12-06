package main

import "fmt"

func main() {
	x := min(2, 3, 5, 2, 1, 0)
	fmt.Printf("The minimum is %d\n", x)
	slice := []int{3, 5, 6, 13, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum is %d\n", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if min < v {
			min = v
		}
	}
	return min
}
