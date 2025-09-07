package main

import "fmt"

func getFullName() (string, string) {
	return "Fitria", "Melliyanni"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}
