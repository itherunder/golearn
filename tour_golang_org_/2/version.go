package test1

import (
	"fmt"
	"runtime"
)

func test2() {
	fmt.Printf("%s\n", runtime.Version())
	fmt.Printf("%+v\n", runtime.Version())
	fmt.Printf("%#v\n", runtime.Version())
	fmt.Printf("%T\n", runtime.Version())
}
