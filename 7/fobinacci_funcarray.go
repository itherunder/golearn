package main

import "fmt"

func main() {
	fmt.Println(fib_slice(5))
}

func fib_slice(n int) []int {
	ans := make([]int, n)
	ans[0], ans[1] = 1, 1
	for i := 2; i < n; i++ {
		ans[i] = ans[i-1] + ans[i-2]
	}
	return ans
}
