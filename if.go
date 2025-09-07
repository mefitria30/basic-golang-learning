package main

import "fmt"

func main() {
	name := "Wonwoo"

	if name == "Fitria" {
		fmt.Println("Hello Fitria")
	} else if name == "Wonwoo" {
		fmt.Println("Hello Wonwoo")
	} else {
		fmt.Println("hi, boleh kenalan?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
