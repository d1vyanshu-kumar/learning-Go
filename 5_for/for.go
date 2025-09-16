package main

//  in go there is not while loop but you can go through for loop and made it like while loop

func main() {
	// for init; condition; post {}
	for i := 0; i < 10; i++ {
		println(i)
	}

	j := 1
	for j < 10 {
		println(j)
		j++
	}

	for k := 0; k < 10; k++ {
		if k%2 == 0 {
			continue
		}
		println(k)
	} // for condition {}



	

	// here  is some new conceot

	for i:= range []int{1, 2, 3, 4, 5} {
		println(i)
	}
	// range is < of that last number.

		// infinite loop

	for {
		println("infinite loop")
	}
}
