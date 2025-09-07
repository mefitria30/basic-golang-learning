package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke ", counter)
		counter++
	}

	fmt.Println("selesai")

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke", counter)
	}

	// akses array
	names := []string{
		"fitria",
		"wonwwo",
		"mingyu",
	}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	namesNew := []string{
		"fitria",
		"wonwwo",
		"mingyu",
	}

	for index, namesNew := range namesNew {
		fmt.Println("index", index, "=", namesNew)
	}

	//contoh lain tanpa index

	for _, namesNew := range namesNew {
		fmt.Println(namesNew)
	}
}
