package main

import "fmt"

// number sequence of a specific length with a a same element
func main() {

	var nums [5]int

	fmt.Println(len(nums)) // length of the array using len() function

	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	fmt.Println(nums) // print the entire array

	fmt.Println(nums[0]) // print the first element of the array

	var vals [4]bool
	fmt.Println(vals) // default value of bool is false
	// [false false false false]

	var names [3]string
	names[0] = "Alice"

	fmt.Println(names) // [Alice  ]

	lums := [5]int{1, 2, 3, 4, 5}

	fmt.Println(lums) // array literal

	// 2d arrays:

	chums := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(chums)

	// if you know the number of elements or fixed size then use arrays cause it a memory efficient
	//  constant time access
	// for dynamic size use slices in GO

}
