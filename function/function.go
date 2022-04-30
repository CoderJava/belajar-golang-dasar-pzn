package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Yudi", "Setiawan")
}

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}
