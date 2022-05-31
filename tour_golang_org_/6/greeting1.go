package test6

import "fmt"

func test6() {
	Greeting("hello: ", "lcvb", "spxp", "dyxccs")
}

func Greeting(prefix string, who ...string) {
	fmt.Printf("%s", prefix)
	for _, v := range who {
		fmt.Printf("%s,", v)
	}
	fmt.Println()
}
