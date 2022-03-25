package test6

func test5() {
	println("In main before calling greeting")
	greeting()
	println("In main after calling greeting")
}

func greeting() {
	println("In greeting: Hi!!!!!!!")
}
