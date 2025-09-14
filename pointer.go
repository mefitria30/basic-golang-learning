package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	/**
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1 // copy by value

	address2.City = "Bandung"
	fmt.Println(address1) // {Jakarta DKI Jakarta Indonesia}
	fmt.Println(address2) // {Bandung DKI Jakarta Indonesia}
	*/

	// pointer yang benar
	var address1 Address = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address2 *Address = &address1
	address2.City = "Bandung"
	fmt.Println(address1) // {Bandung DKI Jakarta Indonesia}
	fmt.Println(address2) // &{Bandung DKI Jakarta Indonesia}
}
