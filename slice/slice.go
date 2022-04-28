package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)      // [Mei Juni Juli]
	fmt.Println(len(slice1)) // 3
	fmt.Println(cap(slice1)) // 8

	months[5] = "Juni edit"
	fmt.Println(slice1) // [Mei Juni edit Juli]

	slice1[0] = "Mei edit"
	fmt.Println(months) // [Januari Februari Maret April Mei edit Juni edit Juli Agustus September Oktober November Desember]
	fmt.Println()

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(days) // [Senin Selasa Rabu Kamis Jumat Sabtu Baru Minggu Baru]
	fmt.Println()

	daySlice2 := append(daySlice1, "Libur Baru") // Karena daySlice1 sudah penuh mmaka, akan membuat array baru dan tidak ada referensi ke daySlice1
	daySlice2[0] = "Sabtu Edit"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(cap(daySlice2)) // Kapasitas daySlice2 masih ada kosong satu
	fmt.Println()

	daySlice3 := append(daySlice2, "Hari Raya") // Karena daySlice2 masih ada yang kosong maka, tidak membuat array baru dan akan referensi ke daySlice2
	daySlice3[0] = "daySlice3"
	fmt.Println(daySlice2)
	fmt.Println(daySlice3)
	fmt.Println()

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Eko"
	newSlice[1] = "Kurniawan"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println((cap(newSlice)))
	fmt.Println()

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println()

	iniArray := [5]int{1, 2, 3, 4, 5}    // Cuma sanggup 5
	iniArray2 := [...]int{1, 2, 3, 4, 5} // Walaupun tidak didefine daya tampung-nya tapi, aslinya cuma sanggup 5. Karena sudah di-inisialization dengan 5 data
	iniSlice := []int{1, 2, 3, 4, 5}     // length dan capability-nya adalah 5 tapi, ini bisa lebih dinamis

	fmt.Println(iniArray)
	fmt.Println(iniArray2)
	fmt.Println(iniSlice)
	iniSlice = append(iniSlice, 6)       // Ini contoh dinamis dari slice. Di array tidak ada fungsi append.
	fmt.Println(iniSlice)
}
