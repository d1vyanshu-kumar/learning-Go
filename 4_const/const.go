package main 

import "fmt"

// name := "John Doe" can't declare outside of a function but you can do with the const keyword "this is only for with the short cut operator"
var name string = "John Doe"

const love = "I love Go!"

func main ()  {
	const Pi = 3.14
	const Greeting = "Hello, World!"

	fmt.Println(Pi)
	fmt.Println(Greeting)
	fmt.Println(love)
	fmt.Println(name)

	// reassigning a constant will cause a compile-time error
	// the meaning of const is that the value cannot be changed

	const (
		StatusOK = 200
		StatusNotFound = 404
		StatusInternalServerError = 500
	)

	fmt.Println(StatusOK, StatusNotFound, StatusInternalServerError)

}