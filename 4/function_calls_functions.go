package main

var a string

func main() {
	a = "G"
	print(a) //"G"
	f1()
}

func f1() {
	a := "O"
	print(a) //"O"
	f2()
}

func f2() {
	print(a) //"G"
}
