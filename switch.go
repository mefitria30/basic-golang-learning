package main

import "fmt"

func main() {
	name := "mingyu"

	switch name {
	case "Fitria":
		fmt.Println("Hello Fitria")
	case "Wonwoo":
		fmt.Println("Hello Wonwoo")
	default:
		fmt.Println("Hi, boleh kenalan?")
	}

	// short statement version
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}