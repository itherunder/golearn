package panic_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

func A1() {
	colorlog.Info("here is A1")
	panic("panic A1")
}

func A2() {
	colorlog.Info("here is A2")
}

func TestPanic(t *testing.T) {
	defer A1() // before panic execute, defer should be executed, as package defer_ said
	// _defer stuct contains the panic struct pointer which panic happened when the _defer execute
	defer A2()
	panic("panic TestPanic")
}

func TestB(t *testing.T) {
	defer B1()
	defer B2()
	//...
	panic("Panic TestB")
	//...
}

func B1() {
	panic("Panic B1")
}

func B2() {
	p := recover() // recover will recover to the func's `deferret` which panic happened
	colorlog.Info("recover p is %v", p)
	panic("Panic B2")
}
