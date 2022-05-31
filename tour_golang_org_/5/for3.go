package test5

import "fmt"

func test6() {
	var i int = 5
	for {
		i--
		fmt.Printf("The variable i is now: %d\n", i)
		if i == 0 {
			break
		}
	}
}
