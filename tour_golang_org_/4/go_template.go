package test4

import "fmt"

const c_ = "C"

var v int = 5

type T struct{}

func init() {

}

func test5() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() {
	//...
}
