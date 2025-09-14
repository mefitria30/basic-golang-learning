package main

import "fmt"

func ups() interface{} {
	// return 1
	return true
	// return "ups"
}

func main() {
	var kosong any = ups()
	fmt.Println(kosong)
}
