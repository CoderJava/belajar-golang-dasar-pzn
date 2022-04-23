package main

import "fmt"

func main() {
	var name string
	var age = 28
	country := "Indonesia" // cuma untuk deklarasi awal saja

	name = "Yudi Setiawan"
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(country)

	name = "CoderJava"
	country = "Singapura"
	fmt.Println(name)
	fmt.Println(country)

	// deklarasi multiple variable yang langsung di-inisialisasikan
	var (
		firstName = "Coder"
		lastName  = "Kotlin"
		salary    = 5000000
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(salary)
}
