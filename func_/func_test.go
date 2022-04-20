package func_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

type Person struct {
	Name, Birth string
	Id          int64
}

func (person *Person) changeName(name string) {
	person.Name = name
}

func (person Person) printMess() {
	colorlog.Info("My name is %v, and my birthday is %v, and my id is %v\n", person.Name, person.Birth, person.Id)
	// person.Id = 2334 // doesn't work
}

func TestFunc(t *testing.T) {
	person := Person{
		"itherunder",
		"1998-01-01",
		2333,
	}
	person.printMess()
	person.changeName("liaozhou")
	person.printMess()
}
