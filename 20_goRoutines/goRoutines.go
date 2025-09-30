package main

import (
	"fmt"
	"sync"
)

func task(id int, w * sync.WaitGroup) {

	defer w.Done() // this is used to tell the wait group that the goroutine has finished its work, and defer let that executed in th end after the function is executed.
	fmt.Printf("Task %d is starting\n", id)

}

func main() {

	var wg sync.WaitGroup // this is used to wait for a collection of goroutines to finish

	for i := 0; i < 10; i++ {

		wg.Add(1) // this is used to tell the wait group that we are going to wait for 1 goroutine to finish
		go task(i , &wg) //  now what we want is that the task should run parallelly there fore wr are going to use "go" keyword 

		// go func(i int) { // this is an anonymous function
		// 	fmt.Println("Hello from anonymous function", i)
		// }(i) // calling the anonymous function immediately and  receving the value of i as parameter
	}

	// this is done bcz the the task function is not executed right away it is scheduled to run and the main function exits before the task function can complete its execution
	// time.Sleep(time.Second * 2) // this is just to make sure that the main function waits for 1 second before exiting so that all the go routines can complete their execution
	// fmt.Println("All tasks are done")
	
	// look in real world application we dont know how much time the task function takes to complete its execution  and in this situationn we are going to use "sync.WaitGroup"

	wg.Wait() // this is used to wait for all the goroutines to finish


}

// in output there is no particular order of execution of the tasks bcz they are running in parallel

// add -> done -> wait