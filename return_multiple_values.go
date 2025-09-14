package main

import "fmt"

func getFullName() (string, string) {
	return "Fitria", "Melliyanni"
}

func main() {
	// jika ingin menampung semua nilai return
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// jika hanya ingin menampung salah satu nilai return
	firstName, _ := getFullName()
	fmt.Println(firstName)
}
