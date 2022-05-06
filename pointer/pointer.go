package main

import "fmt"

func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 // pass by value

	address2.City = "Bandung"

	fmt.Println(address1) // address1 tidak berubah. Kalau di Java, PHP ini pasti akan berubah
	fmt.Println(address2)

	address3 := &address1 // pass by reference (pointer)
	address3.City = "Medan"
	fmt.Println()
	fmt.Println(address1)
	fmt.Println(address3)

	address4 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address5 := &address4
	address6 := &address4
	address7 := &address4

	address5.City = "Aceh"

	address5 = &Address{"Malang", "Jawa Timur", "Indonesia"} // buat pointer baru
	fmt.Println()
	fmt.Println(address4) // tidak berubah ke Malang. Tetap di Aceh
	fmt.Println(address5)

	address6.City = "Palembang"
	*address6 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // paksa reference mengikuti dengan perubahan ini
	fmt.Println()
	fmt.Println(address4) // berubah menjadi Jakarta. Karena address6 awalnya reference ke address4.
	fmt.Println(address5) // tetap di Malang. Karena dia buat pointer baru.
	fmt.Println(address6)
	fmt.Println(address7) // mengikuti perubahan di address6.

	// Cara buat pointer baru dengan data kosong
	alamat := new(Address)
	fmt.Println()
	fmt.Println(alamat)
	alamat.City = "Riau"
	fmt.Println(alamat)
}

type Address struct {
	City     string
	Province string
	Country  string
}
