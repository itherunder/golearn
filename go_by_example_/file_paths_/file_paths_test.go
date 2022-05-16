package file_paths_

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

var pln = fmt.Println

func TestFilePaths(t *testing.T) {
	p := filepath.Join("dir1", "dir2", "file.txt")
	pln("p:", p)

	pln(filepath.Join("dir1//", "file.txt"))
	pln(filepath.Join("dir1//../dir1", "file.txt"))

	pln("Dir(p):", filepath.Dir(p))
	pln("Base(p):", filepath.Base(p))

	pln(filepath.IsAbs("dir/file"))
	pln(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	pln("Ext(filename):", ext)
	pln(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	pln(rel)
}
