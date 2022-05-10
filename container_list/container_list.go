package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Eko")
	data.PushBack("Kurniawan")
	data.PushBack("Khannedy")
	// data.PushFront("Budi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println()

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println()

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}

	fmt.Println()
	fmt.Println(data.Len())
	fmt.Println(data.Front().Value)
	fmt.Println()

	data.Remove(data.Front())
	fmt.Println(data.Front().Value)
}
