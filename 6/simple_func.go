package test6

import "fmt"

func test10() {
	fmt.Printf("The sum of 1,3,5,7,1 is: %d\n", sum(1, 3, 5, 7, 1))
}

func sum(nums ...int) int {
	var ans int = 0
	for _, v := range nums {
		ans += v
	}
	return ans
}
