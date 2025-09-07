package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Fitria"
	names[1] = "Melliyanni"
	names[2] = "Programmer"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(names)

	// bisa juga ditulis gini:
	var values = [3]int{
		98,
		88,
		95,
	}

	fmt.Println(values)

	// bisa diassign kayak gini juga:
	var valuesTest = [...]int{
		100,
		200,
		400,
		78,
	}

	fmt.Println(valuesTest)
	fmt.Println(len(valuesTest))
}
