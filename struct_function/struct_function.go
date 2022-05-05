package main

import "fmt"

func main() {
	customer := Customer{
		Name:    "Yudi",
		Address: "Medan",
		Age:     28,
	}
	// sayHello(customer, "Ditta")
	customer.sayHi("Ditta")
}

type Customer struct {
	Name    string
	Address string
	Age     int
}

func sayHello(customer Customer, name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

// struct function atau struct method
func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}
