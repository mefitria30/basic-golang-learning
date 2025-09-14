package helper

import "fmt"

var version = "1.0.0"      // tidak bisa diakses di luar package helper
var Application = "golang" // bisa diakses di luar package helper

func sayGoodBye(name string) string {
	return "Good Bye " + name
} // tidak bisa diakses di luar package helper

func SayHello(name string) string {
	return "Hello " + name
} // bisa diakses di luar package helper

func Contoh() {
	sayGoodBye("Fitria")
	fmt.Println(version)
}
