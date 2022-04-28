package main

import "fmt"

func main() {
	name := "Yudi"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Yudi" {
		fmt.Println("Hello Yudi")
	} else {
		fmt.Println("Hai, boleh kenalan?")
	}

	// panjangKarakter := len(name)
	// if panjangKarakter > 5 {
	// 	fmt.Println("Panjang karakter nama lebih dari 5")
	// } else {
	// 	fmt.Println("Panjang karakter nama kurang dari atau sama dengan 5")
	// }

	// if short statement (deklarasi variable bisa didalam if)
	if panjangKarakter := len(name); panjangKarakter > 5 {
		fmt.Println("Panjang karakter nama lebih dari 5")
	} else if panjangKarakter == 3 {
		fmt.Println("Panjang karakter nama sama dengan 3")
	} else {
		fmt.Println("Panjang karakter nama kurang dari atau sama dengan 5")
	}
}
