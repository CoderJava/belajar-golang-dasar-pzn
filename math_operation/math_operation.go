package main

import "fmt"

func main() {
	var a = 5
	var b = 4

	var sum = a + b
	fmt.Println(sum)

	var c = a * b
	fmt.Println(c)

	a += 5 // a = a + 5
	fmt.Println(a)

	var d = 10 % 3
	fmt.Println(d)

	b++ // b = b + 1
	fmt.Println(b)

	var negative = -100
	var positive = 100
	fmt.Println(negative)
	fmt.Println(positive)

	var tidakBenar = !true
	fmt.Println(tidakBenar)
}
