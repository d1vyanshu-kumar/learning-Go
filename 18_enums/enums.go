package main


// enumrated types

type orderStatus int

const (
	Received orderStatus = iota
	Processing
	Shipped
	Delivered
)

func changeOrderStatus(os orderStatus) {

	println("Order status changed to", os)
}
func main() {
	changeOrderStatus(Shipped)
}
