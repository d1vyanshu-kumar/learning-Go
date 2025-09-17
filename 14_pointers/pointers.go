package main 

import "fmt"

// whatever the var is you are storing in the memory that var has a address in the memory
// and that address is called pointer

// we are passing by value here
func outer(num int) {
	num = 10
	fmt.Println("Address of num in outer:", &num) // & is used to get the address of the var
}


// this is the way how we can use pointer in go lang and change the original value of the var 

func byRef(num *int) { // here we are using * before the type to indicate that this is a pointer to an int
	*num = 20 // here we are using * before the var to indicate that we want to change the value at the address stored in the pointer
	fmt.Println("Address of num in byRef:", num) // num is a pointer so it will print the address of the var
	fmt.Println("Value of num in byRef:", *num) // here we are using * before the var to get the value at the address stored in the pointer
}

func main() {
	num := 5 // its a copy  we are changing the copy of num not the original num
	fmt.Println("Address of num in main:", &num) // & is used to get the address of the var

	outer(num)
	fmt.Println("Value of num in main after outer call:", num) // num is still 5 because we are passing the value of num to the function not the address of num


	byRef(&num) // here we are using & before the var to get the address of the var
	fmt.Println("Value of num in main after byRef call:", num) // num is now 20 because we are passing the address of num to the function
}