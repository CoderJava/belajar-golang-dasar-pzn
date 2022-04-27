package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Nol"
	names[1] = "Satu"
	names[2] = "Dua"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(len(names))

	var values = [3]int{
		90,
		80,
		75,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))
}
