package reflect_

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/yezihack/colorlog"
)

type Person interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name  string
	Age   int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	colorlog.Info("Hello " + name + ", my name is " + hero.Name)
}

func (hero *Hero) Run() string {
	return fmt.Sprintf("%v run %v km/h", hero.Name, hero.Speed)
}

func TestReflect(t *testing.T) {
	var hero Person
	hero = &Hero{
		"itherunder",
		24,
		12,
	}
	hero.SayHello("liaozhou")
	colorlog.Info(hero.Run())
	typeOfPtrHero := reflect.TypeOf(&Hero{})
	colorlog.Info("Hero's type is %s, kind is %s", typeOfPtrHero, typeOfPtrHero.Kind())
	typeOfHero := typeOfPtrHero.Elem()
	colorlog.Info("Hero's type is %s, kind is %s", typeOfHero, typeOfHero.Kind())
	for i := 0; i < typeOfHero.NumField(); i++ {
		colorlog.Info("field name is %s, type is %s, kind is %s\n",
			typeOfHero.Field(i).Name,
			typeOfHero.Field(i).Type,
			typeOfHero.Field(i).Type.Kind())
	}
	nameField, _ := typeOfHero.FieldByName("Name")
	colorlog.Info("field name is %s, type is %s, kind is %s\n",
		nameField.Name,
		nameField.Type,
		nameField.Type.Kind())

	for i := 0; i < typeOfPtrHero.NumMethod(); i++ {
		colorlog.Info("method name is %s, type is %s, kind is %s\n",
			typeOfPtrHero.Method(i).Name,
			typeOfPtrHero.Method(i).Type,
			typeOfPtrHero.Method(i).Func.Kind())
	}
	sayHelloMethod, _ := typeOfPtrHero.MethodByName("SayHello")
	colorlog.Info("method name is %s, type is %s, kind is %s\n",
		sayHelloMethod.Name,
		sayHelloMethod.Type,
		sayHelloMethod.Func.Kind())
}
