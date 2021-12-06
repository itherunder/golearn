package main

import "fmt"

func main() {
	x2, x3 := getX2AndX3_2(2)
	fmt.Printf("2 * 2 = %d, 2 * 3 = %d\n", x2, x3)
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getX2AndX3_2(input int) (x2, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}
