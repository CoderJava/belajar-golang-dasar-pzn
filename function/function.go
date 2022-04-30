package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Yudi", "Setiawan")
	fmt.Println(getHello("Ditta"))
}

func sayHello() {
	fmt.Println("Hello World")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

