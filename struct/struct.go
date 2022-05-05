package main

import "fmt"

func main() {
	// Cara 1 (Semua field tidak wajib diisi)
	var customer CustomerPremium
	customer.FirstName = "Yudi"
	customer.LastName = "Setiawan"
	customer.Address = "Medan"
	customer.Age = 28
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	// Cara 2 (Semua field tidak wajib diisi)
	customer2 := CustomerPremium{
		FirstName: "Dian",
		LastName:  "Ika Wahyuni",
		Address:   "Jakarta",
		Age:       35,
	}
	fmt.Println(customer2)

	// Cara 3 (Semua field wajib diisi)
	customer3 := CustomerPremium{
		"Ditta",
		"Amelia",
		"Tangerang",
		29,
	}
	fmt.Println(customer3)
}

// Pakai style UpperCamelCase
type CustomerPremium struct {
	FirstName, LastName string
	Address             string
	Age                 int
	// isMarried           bool
}
