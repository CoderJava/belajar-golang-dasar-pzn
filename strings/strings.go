package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan", "Budi"))
	fmt.Println()

	fmt.Println(strings.Split("Eko Kurniawan Khannedy", " "))
	fmt.Println()

	fmt.Println(strings.ToLower("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Khannedy"))
	fmt.Println()

	fmt.Println(strings.Trim("   Eko Kurniawan   ", " "))
	fmt.Println()

	fmt.Println(strings.ReplaceAll("Eko Eko Eko Eko Kurniawan", "Eko", "Budi"))
	fmt.Println()
}
