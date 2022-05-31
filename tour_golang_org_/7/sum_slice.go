package test7

import "fmt"

func test4() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(arr[:]))
}

func sum(ar []int) int {
	s := 0
	for _, v := range ar {
		s += v
	}
	return s
}
