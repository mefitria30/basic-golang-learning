package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	/** 
	var resultString string = result.(string)
	fmt.Println(resultString)

	var resultInt int = result.(int) // panic: interface conversion: interface {} is string, not int
	fmt.Println(resultInt)
	*/

	// type assertion switch
	switch value := result.(type) {
	case string:
		fmt.Println("string:", value)
	case int:
		fmt.Println("int:", value)
	default:
		fmt.Println("unknown type")
	}
}
