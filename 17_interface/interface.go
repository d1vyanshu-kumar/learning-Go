package main

import "fmt"

// An interface is a type that defines a set of method signatures (behavior) that a struct must implement to be considered of that interface type.
// Interfaces allow for polymorphism, enabling different types to be treated uniformly based on shared behavior rather than specific implementations.

// lets start through some real world example make things use case of payment gatway interface.

// 1.st for the payment gatway interface we need a data structure to hold the payment information

type paymenter interface {  // paymenter is the naming convention for interface we use er at the end of the name
	// here we are writing the method signature that we want to implement in our structs
	pay(amount float32) // this is the method signature and if this return you better handle that return type also

	// in go we don't explicitly declare that a struct implements an interface, instead a struct implements an interface by implementing all the methods of that interface by using the method signature.
}

type Payment struct {

	// gatway stripe  //  here we are using the stripe struct to make a payment and this is hardcoded which wil cause issue on unit testing and also if we want to change the payment gatway we need to change the code here so to avoid this we will use interface this will solve the issue of hardcoding and also will make our code more flexible and testable. so the real solution start from interface

	// we dont need to to concrite type things like we did just above and its config.instted we provide the interface

	gateway paymenter

}



// lets make a func and connect that with the payment struct
// you could use this menthod to make a payment as per user.
func (p Payment) makePayment(amount float64) {

	// paymentGateway := razorpay{} // here we are using the razorpay struct to make a payment
	// paymentGateway.pay(float32(amount)) // here we are calling the pay method of the razorpay struct

	// paymentGateway1 := stripe{} // here we are using the stripe struct to make a payment
	// paymentGateway1.pay(float32(amount)) // here we are calling the pay method of the stripe struct

	// here is the best way to do this using structs 

	p.gatway.pay(float32(amount)) // here we are calling the pay method of the stripe struct
}

// add a payment gatway interface

type razorpay struct {}

func (r razorpay) pay (amount float32) {
 // logic to make a payment using Razorpay we can use of the razorpay api here and so on 
 fmt.Println("Payment of", amount, "made using Razorpay")
}

// what if we want to add another payment gatway like strype

// type stripe struct {}

// func (s stripe) pay (amount float32) {
//  // logic to make a payment using Stripe we can use of the stripe api here and so on 
//  fmt.Println("Payment of", amount, "made using Stripe")
// }

func main() {


    // paymentGateway1 := stripe{} 
	// newPayment := Payment{gatway: paymentGateway1 }
	paymentGateway := razorpay{}
	newPayment := Payment{gateway: paymentGateway} // here we are passing the interface to the struct this is called dependency injection and this will make our code more flexible and testable.
	newPayment.makePayment(100.50)
}