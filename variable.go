package main

import "fmt"

func main() {
	var name string

	// bisa juga bentuknya gini:
	// var name = "Fitria"

	name = "Fitria"
	fmt.Println(name)

	// bisa juga bentuknya gini:
	// name := "fitria"
	// tapi ini cuman bisa saat assign pertama

	name = "Fitria Melliyanni"
	fmt.Println(name)

	var (
		firstName = "Fitria"
		lastName  = "Melliyanni"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

}
