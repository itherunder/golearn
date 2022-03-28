package struct_

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yezihack/colorlog"
)

type A struct {
	name string
}

func (a A) Name() string {
	a.name = "Hi! " + a.name
	return a.name
}

func NameOfA(a A) string {
	a.name = "Hi! " + a.name
	return a.name
}

func (pa *A) Name_() string {
	pa.name = "Hi! " + pa.name
	return pa.name
}

func TestStructFunc(t *testing.T) {
	a := A{"eggo"}
	fmt.Println(a.Name()) // same as next line
	fmt.Println(A.Name(a))
	fmt.Println(a.name) // do not be modified, cause the parameter is value delivery

	pa := &a
	fmt.Println(pa.Name_()) // same as next line and next next line
	fmt.Println(a.Name_())  // grammar candy
	fmt.Println((*A).Name_(pa))
	fmt.Println(a.name)

	// t1 == t2 when the func has same parameters type and return type
	t1 := reflect.TypeOf(A.Name)
	t2 := reflect.TypeOf(NameOfA)
	colorlog.Debug("t1 == t2: %v", t1 == t2)

	f1 := A.Name // f1 is a funcval struct pointer!!
	f1(a)

	f2 := a.Name
	f2()
}
