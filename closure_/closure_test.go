package closure

import (
	"fmt"
	"testing"
)

/*
	i被存放到了堆中，其地址被捕获到闭包函数中
	闭包函数的自由变量（即i的地址）通过入口地址的相对+4/8得到
	因此两次输出i都为 2
*/
func create1() (fs [2]func()) {
	for i := 0; i < 2; i++ {
		fs[i] = func() {
			fmt.Println(i)
		}
	}
	return
}

func TestClosure1(t *testing.T) {
	fs := create1()
	for i := 0; i < len(fs); i++ {
		fs[i]() // 2,2
	}
}

var str string = "2333"

func create2() func() {
	return func() {
		str = "hello closure"
	}
}

func TestClosure2(t *testing.T) {
	fs := create2()
	fs()
	fmt.Println(str)
}
