package main

// Tanda _ dipakai jika kita cuma ingin memanggil func init() saja tanpa ada menggunakan property / function yang ada didalam package lain.
// Metode ini disebut dengan blank identifier.

import (
	"belajar-golang-dasar/database"
	"fmt"
	// _ "belajar-golang-dasar/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
