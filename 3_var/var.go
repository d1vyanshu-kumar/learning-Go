package main

import "fmt"

func main() {
	// variable declaration
	var a int = 10
	var b float64 = 3.14
	var c bool = true

	d := "Hello, Go!" // shorthand for variable declaration and initialization

	// printing variables
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// in go you have to make sure that the variable is used otherwise it will throw an error

	// alike in java you write or declare var like this int a = 10;
	// but in go you write var a int = 10;
	// also you can use shorthand declaration like d := "Hello, Go!"

	gh := 100.23
	fmt.Printf("%.1f\n", gh)
}