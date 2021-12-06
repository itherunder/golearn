package main

const Pi = 3.14159265

// const b string = "abc"
const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
const Log2E = 1 / Ln2 // this is a precise reciprocal
const Billion = 1e9   // float constant
const hardEight = (1 << 100) >> 97

// const beef, two, c = "eat", 2, "veg"

// const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

const (
	a = iota //第一个是0，后面的是1，2，后面的不加= iota也可以
	b = iota
	c = iota
)

type Color int

const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	GREEN               // ..
	BLUE
	INDIGO
	VIOLET // 6
)

// const c = getNumber()//error

func getNumber() int {
	return 1
}

func main() {

}
