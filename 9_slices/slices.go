package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array

func main () {
	// unintialized slice is nill(null slice)
	 var nums []int // declaring a slice but there is no size mentioned
	 fmt.Println(nums == nil) // true

	 nums = append(nums, 10) // appending an element to the slice
	 nums = append(nums, 20, 30, 40, 50) // appending multiple elements to the slice

	 fmt.Println(nums)
	 fmt.Println(len(nums)) // length of the slice


	 lums := make([]int, 5) // make function to create a slice with length 5
	 fmt.Println(lums) // [0 0 0 0 0] default value of int is 0
	 fmt.Println(cap(lums)) // capacity of the slice
	 // capacity -> maximum number of elements can fit and this is dyncamic

	//  initial capicity using third parameter of make function
	 chums := make([]int, 3, 5) // length 3 and capacity 5, user 0 insted of 3 to create slice with non zero default values
	 fmt.Println(chums) // [0 0 0]
	 fmt.Println(len(chums)) // 3
	 fmt.Println(cap(chums)) // 5

	 chums[0] = 10
	 chums[1] = 20
	 chums[2] = 30
	 fmt.Println(chums) // [10 20 30]



	 chums = append(chums, 10, 20) // adding two elements to the slice
	 fmt.Println(chums) // [0 0 0 10 20]
	 fmt.Println(len(chums)) // 5
	 fmt.Println(cap(chums)) // 5
	 
	 chums = append(chums, 30) // adding one more element to the slice
	 fmt.Println(chums) // [0 0 0 10 20 30]
	 fmt.Println(len(chums)) // 6
	 fmt.Println(cap(chums)) // 10 -> capacity is doubled

	 tums := [] int {}
	 fmt.Println(tums) // []
	 fmt.Println(len(tums)) // 0
	 fmt.Println(cap(tums)) // 0

	 tums = append(tums, 10)
	 fmt.Println(tums) // [10]
	 fmt.Println(len(tums)) // 1
	 fmt.Println(cap(tums)) // 1

	 tums = append(tums, 20)
	 fmt.Println(tums) // [10 20]
	 fmt.Println(len(tums)) // 2
	 fmt.Println(cap(tums)) // 2

	 // for copying a slice here is the way

	 nums1 := []int{1, 2, 3}
	 nums1 = append(nums1, 4, 5) // adding two more elements to nums1
	 nums2 := make([]int, len(nums1)) // creating a slice with the same length as nums1
	
	 copy(nums2, nums1) // copying elements from nums1 to nums2
	 fmt.Println(nums1) // [1 2 3 4 5]
	 fmt.Println(nums2) // [1 2 3 4 5] -

	 
	 // SLICE OPERATORS

	 nums3 := []int{1, 2, 3, 4, 5}
	 fmt.Println(nums3[0:2]) // [1 2] - from index 0 to index 2 (excluding index 2), if you remove 0 it will consider from the start.
	 fmt.Println(nums3[2:]) // [3 4 5] - from index 2 to the end

	 // slice package 

	 s1 := []int{5, 3, 1, 4, 2}
	 s2 := [] int{5, 3, 1, 4, 2}
	 fmt.Println(slices.Equal(s1, s2)) // true - checks if two slices are equal
}