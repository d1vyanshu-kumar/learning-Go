package main

import "fmt"

func  add(nums ... int ) int { // the three dots before int means that this function can take any number of integer arguments of that type this    is  like a spread operator in JavaScript

	// NOTE HERE WE ARE USING THE THREE DOTS BEFORE THE TYPE NOT AFTER THE PARAMETER NAME
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

// and if you want to recived any type of data type then please use interface{}  like this

 func addAll(args ...interface{})  {
	for _, arg := range args {
		fmt.Println(arg)
	}
}
func main () {
	fmt.Println("Hello, World!") // this println  is called variadic function in Go  where you can pass here n number of parameters

	 addAll(1, "hello", 3.14, true)
	sum := add(1, 2, 3, 4, 5)
	fmt.Println("Sum:", sum)

	lums := []int{6, 7, 8, 9, 10}

	lums1 := add(lums...) // here we are using three dots to spread the slice into individual arguments
	// NOTE WE ARE ADDING THE THREE DOTS AFTER  THE PARAMAETERS NOT BEFORE THE PARAMETERS
	fmt.Println("Lums:", lums1)




}