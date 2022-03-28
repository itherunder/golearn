package type_

import (
	"testing"

	"github.com/yezihack/colorlog"
)

type T struct {
	name string
}

func (tt T) F1() {
	colorlog.Info("t.name: %s", tt.name)
}

func TestType(t *testing.T) {
	tt := T{name: "itherunder"}
	tt.F1() // equal to next line
	T.F1(tt)
}
