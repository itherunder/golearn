//创建一个用于统计字节和字符（rune）的程序，并对字符串 asSASA ddd dsjkdsjs dk 进行分析，
// 然后再分析 asSASA ddd dsjkdsjsこん dk，最后解释两者不同的原因（提示：使用 unicode/utf8 包）。
package test4

import "fmt"

func count_characters(s string) int {
	var cnt int = 0
	for _, c := range s {
		fmt.Print(c)
		cnt++
	}
	fmt.Print(" => ")
	return cnt
}

func test3() {
	var s string = "asSASA ddd dsjkdsjs dk"
	fmt.Printf("%s count : %d\n", s, count_characters(s))
	fmt.Printf("%s count : %d\n", s, len(s))
	var s1 string = "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("%s count : %d\n", s1, count_characters(s1))
	fmt.Printf("%s count : %d\n", s1, len(s1))
}
