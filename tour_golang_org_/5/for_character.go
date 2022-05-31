package test5

import "fmt"

func test3() {
	for i := 1; i <= 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("G")
		}
		fmt.Println()
	}
}
