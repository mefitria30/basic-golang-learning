package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var dataNew Customer
	dataNew.Name = "Mingyu"
	dataNew.Address = "Seoul"
	dataNew.Age = 30

	fmt.Println(dataNew)
	fmt.Println(dataNew.Name)

	// struct literal
	wonwoo := Customer{
		Name:    "Wonwoo",
		Address: "Busan",
		Age:     28,
	}
	fmt.Println(wonwoo)

	// struct literal without field names
	jaehyun := Customer{"Jaehyun", "Incheon", 25}
	fmt.Println(jaehyun)

	// panggil dari struct method
	jaehyun.sayHello("Mingyu")
}
