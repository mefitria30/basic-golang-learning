package main

import "fmt"

func main() {
	// cara 1
	// var person map[string]string = map[string]string{}
	// person["name"] = "Fitria"
	// person["address"] = "Tangerang"

	// cara 2
	person := map[string]string{
		"name":    "Fitria",
		"address": "Tangerang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	// buat map
	book := make(map[string]string)
	book["title"] = "Tutorial Golang"
	book["author"] = "Fitria Melliyanni"
	book["wrong"] = "ups"

	fmt.Println(book)

	delete(book, "wrong")

	fmt.Println(book)
}
