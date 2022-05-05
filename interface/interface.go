package main

import "fmt"

func main() {
	var manusia Person
	manusia.Name = "Eko"
	sayHello(manusia)

	cat := Animal{
		Name: "Pablo",
	}
	sayHello(cat)
}

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// Langsung otomatis terdeteksi implement dari interface HasName
func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

// Langsung otomatis terdeteksi implement dari interface HasName
func (animal Animal) GetName() string {
	return animal.Name
}
