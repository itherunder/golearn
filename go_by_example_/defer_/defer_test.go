package defer_

import (
	"fmt"
	"os"
	"testing"

	"github.com/yezihack/colorlog"
)

func TestDefer(t *testing.T) {
	f := createFile("./defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	colorlog.Info("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	colorlog.Info("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	colorlog.Info("closing")
	err := f.Close()
	if err != nil {
		colorlog.Error("error: %v", err)
		os.Exit(1)
	}
}
