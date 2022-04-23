package main

import "fmt"

func main() {
	const firstName = "Yudi"
	const lastName = "Setiawan"

	// Deklarasi multiple
	const (
		age         = 28
		phoneNumber = "08123456789"
		salary      = 5000000 // Tidak error walaupun tidak dipakai
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(phoneNumber)
}
