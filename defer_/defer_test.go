package defer_

import (
	"testing"
	"unsafe"

	"github.com/yezihack/colorlog"
)

func A1(a int) {
	colorlog.Info("parameter a: %v", a)
}

func TestDefer(t *testing.T) {
	a, b := 1, 2
	defer A1(a) // deferproc function will register defer with its parameter(value delivery) to heap, so the result is `parameter a: 1`

	a = a + b
	colorlog.Info("%v, %v", a, b) // 3, 2
}

func TestDefer1(t *testing.T) {
	a, b := 1, 2
	mp := make(map[string]string)
	mp["2333"] = "23334"

	defer func(i int) {
		a += i // defer func has capture var `a`, `a` will be stored to heap, and the local var `a` will be the heap a's address
		// deferproc will save the address of `a` into registration, and value of `b`
		// cause `b` is not been modified so `b` will not escape
		colorlog.Info("defer %v, %v", a, i) // defer 5, 2
	}(b)

	a += b
	b += a                                                                               // change `b` but will not influence the deferproc registration
	colorlog.Debug("address of a is %v, b is %v, mp is %v", &a, &b, unsafe.Pointer(&mp)) // address of a is 0xc000012348, b is 0xc000012350, mp is 0xc000006048, seems like the address is stack address?
	colorlog.Info("TestDefer1 %v, %v", a, b)                                             // TestDefer1 3, 5
}

func B(a int) int {
	a++
	return a
}

func A2(a int) {
	a++
	colorlog.Info("a: %d", a)
}

func TestDefer2(t *testing.T) {
	a := 1
	defer A2(B(a)) // a: 3, do not be involved with `capture var`, so the parameter of `A2` will be constant when deferproc register is `B(1) = 2`
	a++
	colorlog.Info("a: %d", a) // a: 2
}

// TODO: these notes upon is go1.12's defer, the defer struct will be registered to heap, this need be optimized
// 1.13 save defer struct in stack
// 1.14 canceled defer struct which is above, `open coded defer`
/*
	type _defer struct { // 1.12 saved to heap, 1.13 saved to stack(loop defer still saved to heap), 1.14 cancel this struct, and substituted with expansion in func, seems like inline func
		size int32
		started bool
		heap bool
		sp uintptr
		pc uintptr
		fn *funcval
		_panic *_panic
		link *_defer // the link of defer func
	}

*/

// performance test, 1.14 defer quicker but panic slower(cause when panic, golang need scan stack to find these defers not register to link table to execute)
func BenchmarkDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Defer(i)
	}
}

func Defer(i int) (r int) {
	defer func() {
		r -= 1
		r |= r >> 1
		r |= r >> 2
		r |= r >> 4
		r |= r >> 8
		r |= r >> 16
		r |= r >> 32
		r += 1
	}()
	r = i * i
	return
}
