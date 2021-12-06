//strings 和strconv 库学习
package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.HasPrefix(s, prefix string) bool
	// strings.HasSuffix(s, suffix string) bool
	var str string = "this is an example of a string"
	fmt.Printf("T/F? does the string %#v have prefix %#v? \n", str, "th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "th"))

	// strings.Contains(s string, substr string) bool
	fmt.Printf("T/F? does the string %#v have prefix %#v? \n", str, "example ")
	fmt.Printf("%t\n", strings.Contains(str, "example "))

	// strings.Index(s, str string) int
	fmt.Printf("the index of %#v in %#v is %d\n", "example", str, strings.Index(str, "example"))

	// strings.LastIndex(s, str string) int
	fmt.Printf("the last index of %#v in %#v is %d\n", "a", str, strings.LastIndex(str, "a"))     //22
	fmt.Printf("the last index of %#v in %#v is %d\n", "abc", str, strings.LastIndex(str, "abc")) //-1

	// strings.IndexRune(s string, r rune) int
	fmt.Printf("the last index of %#v in %#v is %d\n", '\u0385', str, strings.IndexRune(str, '\u0385')) //-1

	// strings.Replace(s string, old string, new string, n int) string
	fmt.Println(strings.Replace("233", "2", "3", 1))

	// Count 用于计算字符串 str 在字符串 s 中出现的非重叠次数：
	// strings.Count(s string, substr string) int
	fmt.Println(strings.Count("23333", "333"))

	// Repeat 重叠n次
	// strings.Repeat(s string, count int) string
	fmt.Println(strings.Repeat("23333", 2))

	// 去掉字符串前后的空白字符
	fmt.Println(strings.TrimSpace("  alkjdsflj    "))
	fmt.Println(strings.Trim(" flkajj  ", " "))

	// strings.Fields(s string) []string 利用空白作为分隔符将字符串分割为若干块，并返回一个 slice 。如果字符串只包含空白符号，返回一个长度为 0 的 slice
	// strings.Split(s string, sep string) []string
	fmt.Println(strings.Fields("liao zhou is a stupid asshole."))
	fmt.Println(strings.Split("liaozhou is a asshole", " "))

	// strings.Join(elems []string, sep string) string
	// fmt.Println(strings.Join(["2333" "2334" "2335"], ","))
}
