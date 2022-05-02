package main

import "fmt"

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}

func factorialLoop(value int) int {
	result := 1
	for index := value; index > 0; index-- {
		result *= index
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return value
	} else {
		return value * factorialRecursive(value-1)
	}
}
