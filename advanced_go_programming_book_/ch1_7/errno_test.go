package ch1_7

import (
	"fmt"
	"runtime"
	"syscall"
	"testing"
)

func TestErrno(t *testing.T) {
	err := syscall.Chmod("test.txt", 0777)
	if err != nil {
		fmt.Println(err.(syscall.Errno))
	}
}

func TestRecover(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case runtime.Error:
				fmt.Println("runtime error", x)
			case error:
				fmt.Println("common error", x)
			default:
				fmt.Println("other error", x)
			}
		}
	}()

	panic("error")
}
