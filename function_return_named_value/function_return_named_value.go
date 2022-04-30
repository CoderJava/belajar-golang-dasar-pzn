package main

import "fmt"

func main() {
	namaDepan, namaTengah, namaBelakang := getFullName()
	fmt.Println(namaDepan)
	fmt.Println(namaTengah)
	fmt.Println(namaBelakang)
}

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "Kurniawan"
	lastName = "Khannedy"
	return
}
