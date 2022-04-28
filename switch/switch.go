package main

import "fmt"

func main() {
	name := "Albert"

	switch name {
	case "Eko":
		fmt.Println("Hello Eko")
		fmt.Println("Hello Eko")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Boleh kenalan?")
		fmt.Println("Boleh kenalan?")
	}

	// panjangKarakter := len(name)
	// switch panjangKarakter > 5 {
	// case true:
	// 	fmt.Println("Nama terlalu panjang")
	// case false:
	// 	fmt.Println("Nama sudah benar")
	// }

	// Short statement
	switch panjangKarakter := len(name); panjangKarakter > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	// Switch tanpa kondisi
	usia := 28
	switch {
	case usia > 30:
		fmt.Println("Usia > 30")
	case usia > 20:
		fmt.Println("Usia > 20")
	case usia > 10:
		fmt.Println("Usia > 10")
	default:
		fmt.Println("Usia < 10")
	}

}
