//类型别名
package main

import "fmt"

type TZ int
type Rope string

func main() {
	var a, b TZ = 3, 4
	c := a + b
	var s string = "23333"
	fmt.Printf("c has the value: %v\n", c)
	fmt.Printf("s has the value: %v\n", s)
}
