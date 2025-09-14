package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Fitria")
	fmt.Println(result)
	fmt.Println("Fitria")

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error karena tidak bisa diakses di luar package helper
	// fmt.Println(helper.sayGoodBye("Fitria")) // error karena tidak bisa diakses di luar package helper
}
