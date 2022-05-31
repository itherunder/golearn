package test6

import "fmt"

func test1() {
	i1, _, f1 := ThreeValues()
	fmt.Printf("The int: %d, the float: %f\n", i1, f1)
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.2
}
