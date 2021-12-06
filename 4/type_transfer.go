package main

import "fmt"

//不是使用下划线来分割多个名称
func main() {
	// valueOfTypeB = typeB(valueOfTypeA)
	a := 5.0
	b := int(a)
	fmt.Println(b)
}
