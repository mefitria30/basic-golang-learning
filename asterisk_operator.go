package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	// address2 = &Address{"Malang", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} // dereference pointer
	fmt.Println(address1)                                    // {Bandung DKI Jakarta Indonesia}
	fmt.Println(address2)                                    // &{Malang Jawa Timur Indonesia}
}
