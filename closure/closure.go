package main

import "fmt"

func main() {
	name := "Yudi"
	counter := 0

	increment := func() {
		// a := "A" --> Variable ini tidak bisa diakses dari luar anonymous function
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	// a = "edit A"

	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
