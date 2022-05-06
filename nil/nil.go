package main

import "fmt"

func main() {
	person1 := NewMap("Eko")
	fmt.Println(person1)

	person2 := NewMap("")
	if person2 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
