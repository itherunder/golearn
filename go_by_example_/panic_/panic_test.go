package panic_

import (
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	// panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
