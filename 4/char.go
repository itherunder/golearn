//字符类
package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U\n", ch, ch2, ch3) // UTF-8 code point
	fmt.Println(unicode.IsDigit(rune(ch)), unicode.IsLetter(rune(ch)), unicode.IsSpace(rune(ch)))
}
