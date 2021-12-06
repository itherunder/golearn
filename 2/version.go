package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Printf("%+v\n", runtime.Version())
	fmt.Printf("%#v\n", runtime.Version())
	fmt.Printf("%T\n", runtime.Version())
}
