package main

import "fmt"

func main() {
	names := [...]string{
		"wonwoo",
		"dokyeom",
		"scoup",
		"mingyu",
		"vernon",
		"jun",
	}

	slice1 := names[4:6]
	fmt.Println(slice1) // vernon, jun

	slice2 := names[:3]
	fmt.Println(slice2) // wonwoo, dokyeom, scoup

	var slice3 []string = names[3:]
	fmt.Println(slice3) // mingyu, vernon, jun

	var slice4 []string = names[:]
	fmt.Println(slice4) // [wonwoo dokyeom scoup mingyu vernon jun]

	// append slice

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur baru")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// make slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Fitria"
	newSlice[1] = "Fitria"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Fitria")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	// array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
