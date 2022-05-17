package command_line_arguments_

import (
	"fmt"
	"os"
	"testing"
)

func TestCommandLineArguments(t *testing.T) {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
