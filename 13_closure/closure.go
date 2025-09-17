package main 

import "fmt"

func outer() func() int {

	counter := 0

	return func() int {
		counter++
		return counter
	}
}


func main() {

	increment := outer()

	fmt.Println(increment())
	fmt.Println(increment())
}