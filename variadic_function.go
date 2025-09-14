package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	// slice paramater
	numbers := []int{1, 2, 3, 4, 5}
	total = sumAll(numbers...)
	fmt.Println(total)
}
