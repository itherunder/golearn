package exit_

import (
	"fmt"
	"os"
	"testing"
)

func TestExit(t *testing.T) {
	defer fmt.Println("!") // This line is never reached.

	os.Exit(3)
}
