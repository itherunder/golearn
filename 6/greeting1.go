package main

import "fmt"

func main() {
	Greeting("hello: ", "lcvb", "spxp", "dyxccs")
}

func Greeting(prefix string, who ...string) {
	fmt.Printf("%s", prefix)
	for _, v := range who {
		fmt.Printf("%s,", v)
	}
	fmt.Println()
}
