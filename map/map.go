package main

import "fmt"

func main()  {
	person := map[string]string{
		"name": "Yudi",
		"address": "Medan",
	}	
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["title"] = "Programmer"
	fmt.Println(person["title"])
	person["title"] = "Product Engineer"
	fmt.Println(person["title"])
	fmt.Println()
	
	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Eko"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}