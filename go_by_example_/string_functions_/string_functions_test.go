package string_functions_

import (
	s "strings"
	"testing"

	"github.com/yezihack/colorlog"
)

var p = colorlog.Info

func TestStringFunctions(t *testing.T) {
	p("Contains:  %v", s.Contains("test", "es"))
	p("Count:     %v", s.Count("test", "t"))
	p("HasPrefix: %v", s.HasPrefix("test", "te"))
	p("HasSuffix: %v", s.HasSuffix("test", "st"))
	p("Index:     %v", s.Index("test", "e"))
	p("Join:      %v", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    %v", s.Repeat("a", 5))
	p("Replace:   %v", s.Replace("foo", "o", "0", -1))
	p("Replace:   %v", s.Replace("foo", "o", "0", 1))
	p("Split:     %v", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   %v", s.ToLower("TEST"))
	p("ToUpper:   %v", s.ToUpper("test"))

	p("Len:  %v", len("hello"))
	p("Char: %v", "hello"[1])
}
