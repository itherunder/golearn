package test4

var a_ string

func test4() {
	a_ = "G"
	print(a_) //"G"
	f1()
}

func f1() {
	a_ := "O"
	print(a_) //"O"
	f2()
}

func f2() {
	print(a_) //"G"
}
