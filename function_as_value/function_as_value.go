package main

import "fmt"

func main() {
	sayGoodBye := getGoodBye
	result := sayGoodBye("Eko")

	fmt.Println(getGoodBye("Yudi"))
	fmt.Println(sayGoodBye("Pablo"))
	fmt.Println(result)
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}
