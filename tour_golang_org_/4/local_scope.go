package test4

var a___ = "G"

func test7() {
	n() //"G"
	m() //"0"
	n() //"G"
}

func n() { print(a___) }

func m() { a___ := "0"; print(a___) }
