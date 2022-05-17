package command_line_flags_

import (
	"flag"
	"fmt"
	"testing"
)

func TestCommandLineFlags(t *testing.T) {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// testv := flag.Bool("test.v", false, "a bool var")
	// testtimeout := flag.String("test.timeout", "default", "a string var")
	// testrun := flag.String("test.run", "default", "a string var")
	// testcount := flag.Int("test.count", 1, "an int var")

	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
	// fmt.Println("testv:", testv)
	// fmt.Println("testtimeout:", testtimeout)
	// fmt.Println("testrun", testrun)
	// fmt.Println("testcount", testcount)
}
