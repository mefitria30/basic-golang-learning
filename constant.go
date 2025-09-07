package main

import "fmt"

func main() {
	// const firstName string = "Fitria"
	// const lastName = "Melliyanni"

	// atau bentuknya bisa multiple gini:
	const (
		firstName string = "Fitria"
		lastName         = "Melliyanni"
	)
	// error :  cannot assign to lastName (neither addressable nor a map index expression)
	// lastName = "Melliyanni ver 2"

	fmt.Println(firstName)
	fmt.Println(lastName)

}
