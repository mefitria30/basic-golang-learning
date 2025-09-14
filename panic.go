package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// ini contoh recover yang benar karena defer dijalankan ketika panic terjadi
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}

	// ini contoh recover yang salah
	// message := recover()
	// fmt.Println("terjadi panic", message)
}

func main() {
	runApp(true)
	fmt.Println("Aplikasi berjalan")
}
