package helper

import "fmt"

// exported berarti bisa diakses dari luar package
// unexported berarti tidak bisa diakses dari luar package

// var versi = 1 // private (unexported)
var Application = "Belajar Golang" // public (exported)

// public (exported)
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// private (unexported)
// func sayGoodbye(name string) {
// 	fmt.Println("Goodbye", name)
// }
