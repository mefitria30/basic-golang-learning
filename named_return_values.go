package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Fitria"
	middleName = "Wonwoo"
	lastName = "Mingyu"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
