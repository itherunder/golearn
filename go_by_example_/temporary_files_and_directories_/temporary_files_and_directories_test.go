package temporary_files_and_directories_

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestTemporaryFilesAndDirectories(t *testing.T) {
	f, err := os.CreateTemp("", "example")
	check(err)

	fmt.Println("Temp file name: ", f.Name())

	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp("", "exampledir")
	check(err)
	fmt.Println("Temp dir name: ", dname)

	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2, 3, 4}, 0644)
	check(err)
}
