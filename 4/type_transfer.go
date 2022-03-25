package test4

import "fmt"

//不是使用下划线来分割多个名称
func test12() {
	// valueOfTypeB = typeB(valueOfTypeA)
	a := 5.0
	b := int(a)
	fmt.Println(b)
}
