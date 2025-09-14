package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var dataNew Customer
	dataNew.Name = "John"
	dataNew.Address = "New York"
	dataNew.Age = 30

	fmt.Println(dataNew)
}
