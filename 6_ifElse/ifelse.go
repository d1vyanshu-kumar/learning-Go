package main


func main() {

	age :=18

	if age < 18 {
		println("You are a minor.")
	} else if age >= 18 && age < 65 {
		println("You are an adult.")
	} else {
		println("You are a senior.")
	}

	if num := 10; num%2 == 0 {
		println(num, "is even.")
	} else {
		println(num, "is odd.")
		println(age) // you can use age variable here 
	}

	// you can declare variable in if statement and use it in else if and else block
	// but you can't use it outside the if-else block

	// short statement
	// you can use short statement in if-else block
	// but you can't use it outside the if-else block

	// you can use multiple conditions in if-else block 

	//  GO DOESN'T HAVE TERNARY OPERATOR.
}