package main

import "fmt"

func main() {
	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	changeCountryToIndonesia(alamat)
	fmt.Println(alamat) // Property country tidak berubah

	changeCountryToIndonesia2(&alamat)
	fmt.Println(alamat) // Property country berubah karena menggunakan pointer

	alamat2 := &Address{
		City:     "Medan",
		Province: "Sumatera Utara",
		Country:  "",
	}
	changeCountryToIndonesia2(alamat2)
	fmt.Println(alamat2) // Property country berubah karena menggunakan pointer
}

type Address struct {
	City     string
	Province string
	Country  string
}

func changeCountryToIndonesia(address Address) {
	address.Country = "Indonesia"
}

func changeCountryToIndonesia2(address *Address) {
	address.Country = "Indonesia"
}
