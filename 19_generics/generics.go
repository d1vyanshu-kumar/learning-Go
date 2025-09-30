package main

import "fmt"

//[T interface{}]
func printSlice[T int | string ](s []T) { // insted of declaring so many types we can use "comparable" here
	for _, v := range s {
		fmt.Println(v)
	}
}

type stack [T any] struct {
	// LIFO
	items []T
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	printSlice(ints)

	strs := []string{"a", "b", "c", "d"}
	printSlice(strs)

	myStack := stack[int]{
		items: []int{1, 2, 3, 4, 5},
	}

	fmt.Println(myStack)

}
