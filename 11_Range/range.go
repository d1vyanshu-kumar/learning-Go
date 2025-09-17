package main

import "fmt"

// it is used for itrating the data structure
func main () {
	 
	nums := []int{10, 20, 30, 40, 50}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i]) // traditional way of iterating
	}

	// using range keyword
	 sum := 0
	for index, num := range nums {
		fmt.Println("Index:", index) // index of the element
		sum += num
		fmt.Println(num) // range way of iterating
	}
	fmt.Println("Sum:", sum)

	m := map[string]int{
		"Alice":  25,
		"Bob":    30,
		"Charlie":35,
	}

	for key, value := range m {
		fmt.Println("Key:", key) // key of the map
		fmt.Println("Value:", value) // value of the map
	}

	 for index, char := range "hello" {
		fmt.Println("Index:", index) // index of the character
		fmt.Println("Unicode:", char) // Unicode code point of the character
		fmt.Println("Character:", string(char)) // character of the string
	 }
}