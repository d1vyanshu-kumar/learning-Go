package main

import (
	"fmt"
	"maps"
)

func main() {
	// associative type data structure

	// how can we create a map we are going to use make function like we did in slices but here we need to provide its data type of key and value

	user_Map := make(map[int]string) // key is int and value is string

	fmt.Println(user_Map) // map[]

	// adding key value pairs to the map
	user_Map[1] = "Alice"
	user_Map[2] = "Bob"
	user_Map[3] = "Charlie"

	// get an element from the map
	fmt.Println(user_Map[2]) // Bob

	// length of the map
	fmt.Println(len(user_Map)) // 3

	// deleting an element from the map
	delete(user_Map, 2)
	fmt.Println(user_Map) // map[1:Alice 3:Charlie]


	fmt.Println(user_Map[20]) // if key does not exists in the map then it returns zero value you know what is zero means in programming.

	// for clearing the map
	// clear(user_Map);


	///--------------------- we can create map without using make function ------------------///

	 new_Map := map[string]int{
	 	"Apple":  1,
	 	"Banana": 2,
	 	"Orange": 3,
		"key":  0, 
	 }

	 fmt.Println(new_Map) // map[Apple:1 Banana:2 Orange:3]


	 // for checking if a key exists in the map or not here we are going to use this way:

	 v,ok := new_Map["key"]

	 fmt.Println(v) // provide the written value by this way

	 if ok {
		fmt.Println("Key exists in the map")
	 } else {
		fmt.Println("Key does not exist in the map")
	 }


	 // by  this way we can check of the two map are equal or not

	 m1 := map[int]string{
	 	1: "Alice",
	 	2: "Bob",
	 }
	 
	 m2 := map[int]string{
	 	1: "Alice",
	 	3: "Bob",
	 }

	//  fmt.Println(m1 == m2) // invalid operation: m1 == m2 (map can only be compared to nil)

	 fmt.Println(maps.Equal(m1, m2)) // true - checks if two maps are equal
}