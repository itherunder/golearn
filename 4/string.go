//string类型
package main

import "fmt"

func main() {
	var s string = "2333" + "2334" //最好用Bytes.Buffer来做字符串拼接，效率更高
	s += "2335"
	fmt.Printf("%s", s)
}
