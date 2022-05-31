package test4

import (
	"fmt"
	"os"
	"runtime"
)

//这些一般都是用来声明包级别的全局变量
var n_ int = 64
var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)

var str string = "this is string"

func tset15() {
	a, b, c := 5, 7, "abc"
	a, b = b, a
	fmt.Println(a, b, c)
	goos := runtime.GOOS //简短声明赋值
	fmt.Printf("the operation system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("path is: %s\n", path)
}
