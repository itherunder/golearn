package main

import "fmt"

func main() {
	for i := 1; i <= 15; i++ {
		fmt.Println(i)
	}
	cnt := 0
LOOP:
	if cnt >= 15 {
		return
	}
	fmt.Println(cnt)
	cnt++
	goto LOOP
}
