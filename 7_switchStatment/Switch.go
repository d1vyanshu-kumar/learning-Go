package main

import (
	"fmt"
	"time"
)

func main() {


	// there is no break statement in go switch
	i:=5
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Not found")
	}

	// multiple conditions switch
	// isn't this is a cool thing!

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("I'm an int and my value is %d\n", t)
			case string:
			fmt.Printf("I'm a string and my value is %s\n", t)
			case bool:
			fmt.Printf("I'm a bool and my value is %t\n", t)
		default:
			fmt.Printf("Don't know my type %T\n", t)
		}
	}

	whoAmI(42)
	whoAmI("hello")
	whoAmI(true)
	whoAmI(3.14)
}