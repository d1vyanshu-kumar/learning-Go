package main

import (
	"fmt"
	"sync"
)

// mutex- mutual exclusion
// -- when we do multihreading we use to use mutex to prevent race conditions

// race conditions- when multiple prcesses try to modify same resource at the same time will cause of race conditions

// lets took some real exmple lets says we made a some socal media app

type post struct {
	views int


	// here is the ending part how can we solve this race conditions by using mutex: and we are using mutex inside the struct generally
	mu sync.Mutex         // now the 2nd part is go to the inc func and lock that by using mutex -----E01
}

func (p *post) inc(wg *sync.WaitGroup) { // atteched the structs in the func
    defer func ()  {
		p.mu.Unlock() // unlock the mutex to allow other goroutines to enter this section ---------E03
		wg.Done() // mark the goroutine as done when the function returns
	}()
	p.mu.Lock()		 // lock the mutex to prevent other goroutines from entering this section ---------E02
	// important dont put lock in all logic just put there where we modify the data.like here we are modifying the views.
	p.views += 1


	// here we are using mutex to lock the critical section of the code where we are modifying the views field of the post struct.
	// this will prevent other goroutines from entering this section until we unlock the mutex.
	// this will ensure that only one goroutine can modify the views field at a time.
}

func main() {

	var wg sync.WaitGroup // wait group is used to wait for a collection of goroutines to finish

	myPost := post{
		views: 0,
	}

	for i := 0; i < 100; i++ {

		wg.Add(1)          // we are adding one goroutine to the wait group
		go myPost.inc(&wg) // as we know in go we use go if we want to run process concurrently

		// we don't use we.done here cause the above line is not cause blocking it means it execute and move to next line.for this we need to implement inside that function.

		// 		go → launches the function as a goroutine (runs in the background, concurrently).

		// myPost.inc(&wg) → calls the inc method on myPost, passing &wg.

		// &wg → means “the address of wg,” i.e., a pointer to the WaitGroup.
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println(myPost.views)

}

// Why Not Just wg.Add(100) Once?
// You could do that before the loop, but then you must be very careful that you always launch exactly 100 goroutines. By putting wg.Add(1) inside the loop, you’re saying:

// “I’m adding one box only when I actually launch one goroutine.” This keeps the checklist perfectly in sync with the goroutines you start.

// Step-by-Step with 5 Goroutines
// Start: Before launching goroutines, the WaitGroup counter = 0.

// Loop launches 5 goroutines: Each time you call wg.Add(1), the counter goes up by 1. After 5 iterations:

// Code
// Counter = 5
// Now goroutines start finishing: Each goroutine calls wg.Done() when it’s done. wg.Done() is equivalent to wg.Add(-1) → it subtracts 1 from the counter.

// First goroutine finishes → counter goes from 5 → 4

// Second finishes → counter goes from 4 → 3

// Third finishes → counter goes from 3 → 2

// Fourth finishes → counter goes from 2 → 1

// Fifth finishes → counter goes from 1 → 0

// When counter = 0: wg.Wait() unblocks, because the WaitGroup knows all goroutines are done.
