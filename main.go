package main

import (
	"flag"
	"golearn/sha256_"
)

// import "golearn/turingbot"

func testBot() {
	// turingbot.TuringBot()
}

func test4_2() {
	var type_ string
	flag.StringVar(&type_, "t", "SHA256", "type of sha function")
	flag.Parse()
	sha256_.Sha(type_)
}

func main() {
	// fmt.Println("hello world")
	test4_2()
}
