package main

import (
	"fmt"
	// "time"
)

// sending
// func processNum(numChannel chan int) {

// 	// fmt.Println("number received is: ", <- numChannel) // receiving the data from channel

// 	for num := range numChannel {
// 		fmt.Println("number received is: ", num)
// 	}

// }

// func sum(result chan int, num1 int, num2 int) {

// 	result <- (num1 + num2) // sending the data to channel.    "	blocking operation"


// }

// // receiving. gorutine syncronizer
//  func task(result chan bool) {

// 	defer func ()  {
// 		result <- true
// 	}()
// 	fmt.Println("task started")
//  }

func main() {


	// result := make(chan bool)

	// go task(result)

	// <- result // blocking operation until we get the value from the channel


	// result := make(chan int)

	// go sum(result, 5, 6)

	// res := <- result
	
	// fmt.Println("result is: ", res)


	//---------------- channel is always bloking ----------------//
    // here is the syntext first give a name of channel
	// through channel we can share the data b/w the goroutines


	// messageChannel := make(chan string)

	// messageChannel <- "hello channel" // sending the data to channel.    "	blocking operation"

	// // receiving the data from channel

	// msg := <- messageChannel

	// fmt.Println(msg) // all goroutines are asleep - deadlock!

	// ---------part 2 ----------------//

	// numChannel := make(chan int)

	// go processNum(numChannel)
	// // we need to stop for some second to let the goroutine start

	// // numChannel <- 5  // sending the data to channel.    "	blocking operation" and this is unbufferd channel thats why it is blocking and it is very important to undersrand that if you implemetn a queue like system then using this is very inefficient bcz it is blocking operation.

	// for i := 0; i < 11; i++ {
	// 	numChannel <- i

	// }

    // time.Sleep(2 * time.Second)


	// we are using this like a queue in a go using loops and so on ...


	//---- buffered channel -----//

	// look if we are using the unbuffered channel then it means that we are sending one data at a time and until that data is not recieved the next data will not be sent


	// and in buffered channel we can send the limited amount of data without blocking.

	// lets take a exmple we are implementing a queue system for an email system

	// emailChannel := make(chan string, 100) // here 100 is the buffer size means we can send 100 emails without blocking

	// // sync

	// done := make(chan bool)

	// go emailWorker(emailChannel, done)
    

	// emailChannel <- "email 1". //  using loop for this
	// emailChannel <- "email 2"
	// emailChannel <- "email 3"
	// emailChannel <- "email 4"
	// emailChannel <- "email 5"

	// fmt.Println(<- emailChannel)
	// fmt.Println(<- emailChannel)
	// fmt.Println(<- emailChannel)
	// fmt.Println(<- emailChannel)
	// fmt.Println(<- emailChannel)

	// for i := 0; i < 10; i++ {
	// 	emailChannel <- fmt.Sprintf("email %d", i+1)
	// }

	// close(emailChannel) // this is used to close the channel when we are done sending data to it. this is important to avoid deadlock in the receiving goroutine
	// <- done // blocking operation until we get the value from the channel

	// now what if we want to send and receive multiple channels

	channel1 := make(chan int)
	channel2 := make(chan string)

	// sending the data to channel1 by usinf inline goroutine
    // as you can see here 2 go routine here sending data to 2 different channels
	go func() {
		channel1 <- 42
	}()

	go func() {
		channel2 <- "hello"
	}()

	// and if you want to recieved the data from multiple channels we can use here combination  and using "select" keyword
	
	// 1st needed a for loop
	// and here recieving data from multiple channels using select
	for i := 0; i < 2; i++ {
		select {
		case num := <-channel1:
			fmt.Println("Received from channel1:", num)
		case msg := <-channel2:
			fmt.Println("Received from channel2:", msg)
		}
	}

}

// we are sending and receiving a single channel as we can see here...
// func emailWorker(emailChannel chan string, done chan bool)  {   how to make this a recieve only channel as for now you can send data and receive data from same channel but if you are using this "emailChannel <- chan string" this makes its a recieve only channel you can't send data from here like  emailChannel <- "email 1" this will give you an error and what if we make it send only same but direction opp of arrow "done chan <- bool" this will make it send only channel you can't receive data from here like <- done this will give you an error see with respect to "CHAN"

// 	defer func() {
// 		done <- true
// 	}()

// 	for email := range emailChannel {
// 		fmt.Println("Sending email to ", email)
// 		time.Sleep(time.Second)
// 	}

// }




// go routine is a vert important concept if you want to revise then real concept start from 13 -> variadic.