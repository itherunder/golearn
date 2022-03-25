package test4

var a__ = "G"

func test16() {
	n__() //"G"
	m_()  //"O"
	n__() //"O"
}

func n__() {
	print(a__)
}

func m_() {
	a__ = "O"
	print(a__)
}
