package main

import "fmt"

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	fmt.Println(lastName)
}

func getFullName() (string, string, string) {
	return "Eko", "Kurniawan", "Khannedy"
}
