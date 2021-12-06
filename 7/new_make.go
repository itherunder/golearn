// new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体（参见第 10 章）；它相当于 &T{}。
// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）。

package main

import "fmt"

func main() {
	arr := new([]int) // *p == nil; with len and cap 0
	*arr = append(*arr, 2)
	fmt.Println(arr)
	arr1 := make([]int, 10)
	fmt.Println(arr1)
}
