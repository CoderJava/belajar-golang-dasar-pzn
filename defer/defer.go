package main

import "fmt"

func main() {
	runApplication(0)
}

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() // Defer harus diletakkan dipaling atas
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}
