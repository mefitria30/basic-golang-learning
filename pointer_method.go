package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	newData := Man{"Tom"}
	newData.Married()
	fmt.Println(newData.Name) // Output: Tom
}