package main

import "fmt"

func main() {
	n := 0
	Multiply(2, 3, &n)
	fmt.Printf("2 * 3 = %d\n", n)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
