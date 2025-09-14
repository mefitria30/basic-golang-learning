// defer akan menunda eksekusi sebuah fungsi sampai fungsi yang memanggilnya selesai dieksekusi.
// Fungsi yang ditunda akan dieksekusi dalam urutan LIFO (Last In, First Out).

package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // defer akan menunda eksekusi fungsi logging sampai runApplication selesai dieksekusi
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
