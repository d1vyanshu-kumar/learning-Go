package main

import "fmt"

// Structs in Go are a way to group related data together into a single unit. They are similar to classes in other programming languages, but they do not have methods or inheritance. Structs are used to create custom data types that can hold multiple fields of different types.

type Person struct {
	Name string
	Age  int
}

func main() {
	// Creating a new instance of the Person struct
	p := Person{Name: "John", Age: 30}

	// Accessing the fields of the struct
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// Modifying the fields of the struct
	p.Age = 31
	fmt.Println("Updated Age:", p.Age)

	// You can also create a pointer to a struct
	p2 := &Person{Name: "Jane", Age: 25}
	fmt.Println("Name:", p2.Name)
	fmt.Println("Age:", p2.Age)

	// Modifying the fields of the struct through the pointer
	p2.Age = 26
	fmt.Println("Updated Age:", p2.Age)
}