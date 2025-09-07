package main

import "fmt"

func getHello(name string) string {
	hello := "Hello " + name
	return hello
}

func main() {
	result := getHello("Fitria")
	fmt.Println(result)

	fmt.Println(getHello("Wonwoo"))
}
