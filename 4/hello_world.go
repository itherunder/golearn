package main

import (
	fm "fmt"    //别名
	_ "runtime" //这样仅会调用该包的init函数
)

// init is the initialization of this package
func init() {
	fm.Println("init function")
}

// main is the enterency of a package
func main() { //init函数会优先于main
	fm.Println("hello, world")
}
