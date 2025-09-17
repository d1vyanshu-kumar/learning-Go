package main

import "fmt"

// main entry point of the program which is directly called by the go runtime

func addInteger(a int, b int) int { // note here we are going to define which type of value this function is going to return here it is "int"
	return a + b
}

func progLang() (string, bool) { // function returning multiple values
	return "Go", true
}

func fnExample(fn func(a int) int) {
	fn(1)
}

func sybnc() func(a int) int {

	return  func(a int) int {
		return 3
	}
}

func main() {

	ln := sybnc()
	ln(4)

	fn := func(a int) int {
		return 2
	}
	fnExample(fn)

	sum := addInteger(10, 20) // function call
	fmt.Println(sum)          // 30

	lang1, lang2 := progLang() // function call
	fmt.Println(lang1, lang2)  // Go true
}
