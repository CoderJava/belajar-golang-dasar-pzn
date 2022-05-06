package main

import "fmt"

func main() {
	result := random()

	// Type assertion dengan cara tidak aman
	// resultString := result.(string)
	// fmt.Println(resultString)

	// Type assertion dengan cara aman
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

func random() interface{} {
	return true
}
