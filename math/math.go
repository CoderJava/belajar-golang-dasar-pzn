package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))   // Dibulatkan paling dekat
	fmt.Println(math.Round(1.3))   // Dibulatkan paling dekat
	fmt.Println(math.Floor((1.7))) // Dipaksa bulatkan ke bawah
	fmt.Println(math.Ceil(1.3))    // Dipaksa bulatkan ke atas
	fmt.Println(math.Max(10, 20))  // Cari nilai terbesar
	fmt.Println(math.Min(10, 20))  // Cari nilai terendah
}
