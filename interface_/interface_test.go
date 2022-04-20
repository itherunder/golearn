package interface_

import (
	"io"
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

type Tank interface {
	Walk()
	Fire()
}

type Plane interface {
	Fly()
}

type PlaneTank interface {
	Tank
	Plane
}

type Printer interface {
	Print(interface{})
}

type FuncCaller func(p interface{})

func (funcCaller FuncCaller) Print(p interface{}) {
	funcCaller(p)
}

func TestPrinter(t *testing.T) {
	var printer Printer
	printer = FuncCaller(func(p interface{}) {
		colorlog.Info("%v", p)
	})
	printer.Print("liaozhou is itherunder.eth")
}

func TestInterface(t *testing.T) {
	var e interface{}
	f, err := os.Open("./test.txt")
	if err != nil {
		colorlog.Error("open file error: %v", err)
	}
	e = f // assign a eface will assign it's _type to the _type of f
	// _type of a type is global variable, so it's value is the same as the value of f's _type
	colorlog.Info("e type: %T", e)
	r, ok := e.(*os.File) // eface.(type)
	if ok {
		colorlog.Info("r type: %T", r)
	}

	var rw io.ReadWriter
	r, ok = rw.(*os.File) // iface.(type)
	// assign a iface will compare the `itable` of rw and f
	// itable stored in a hash table, the key is <_type, *_type>, like <io.ReadWriter, *os.File>
	if ok {
		colorlog.Info("r type: %T", r)
	}

	// eface.(iface)
	var e2 interface{}
	f2, err := os.Open("./test.txt")
	if err != nil {
		colorlog.Error("open file error: %v", err)
	}
	e2 = f2
	rw2, _ := e2.(io.ReadWriter)
	colorlog.Info("rw2 type: %T", rw2)

	// iface.(iface)
	var w io.Writer
	f3, err := os.Open("./test.txt")
	if err != nil {
		colorlog.Error("open file error: %v", err)
	}
	w = f3 // f3 has Write func, so this assign can success
	rw3, _ := w.(io.ReadWriter)
	colorlog.Info("rw3 type: %T", rw3)
}
