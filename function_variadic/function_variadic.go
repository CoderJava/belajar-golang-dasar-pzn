package main

import "fmt"

func main() {
	name, total := sumAll("Yudi", 10, 10, 10)
	fmt.Println(name, total)

	slice := []int{10, 10, 10, 10, 10}
	fmt.Println(sumAll("Budi", slice...))
}

// Variadic function hanya bisa diletakkan di parameter terakhir atau yang paling kanan
func sumAll(name string, numbers ...int) (string, []int) {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return name, numbers
}
