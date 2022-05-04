package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Nilai x=")
	scanner.Scan()
	x, error := strconv.Atoi(scanner.Text())
	if error != nil {
		fmt.Print("Nilai x tidak valid. Harus angka ya.")
		return
	}

	fmt.Print("Masukkan Nilai y=")
	scanner.Scan()
	y, error := strconv.Atoi(scanner.Text())
	if error != nil {
		fmt.Print("Nilai y tidak valid. Harus angka ya.")
		return
	}

	fmt.Println()
	fmt.Println("Semua akan menghasilkan nilai True atau False")
	fmt.Println("Keterangan: Nilai 1 = True dan Nilai 0 = False")
	fmt.Println("==============================================")

	fmt.Println()
	fmt.Println("Operator Relasi")
	result1 := x == y
	result2 := x != y
	var (
		strResult1 int
		strResult2 int
	)
	if result1 {
		fmt.Println("Hasil dari", x, "==", y, "--> 1")
		strResult1 = 1
	} else {
		fmt.Println("Hasiil dari", x, "==", y, "--> 0")
		strResult1 = 0
	}
	if result2 {
		fmt.Println("Hasil dari", x, "!=", y, "--> 1")
		strResult2 = 1
	} else {
		fmt.Println("Hasil dari", x, "!=", y, "--> 0")
		strResult2 = 0
	}
	fmt.Println()

	fmt.Println("Operator Logika")
	printOutputOperatorLogikaAND(result1, result2, strResult1, strResult2)
	printOutputOperatorLogikaOR(result1, result2, strResult1, strResult2)
	fmt.Println("==============================================")

	fmt.Println()
	fmt.Println("Operator Relasi")
	result1 = x > y
	result2 = x < y
	if result1 {
		fmt.Println("Hasil dari", x, ">", y, "--> 1")
		strResult1 = 1
	} else {
		fmt.Println("Hasil dari", x, ">", y, "--> 0")
		strResult1 = 0
	}
	if result2 {
		fmt.Println("Hasil dari", x, "<", y, "--> 1")
		strResult2 = 1
	} else {
		fmt.Println("Hasil dari", x, "<", y, "--> 0")
		strResult2 = 0
	}
	fmt.Println()

	fmt.Println("Operator Logika")
	printOutputOperatorLogikaAND(result1, result2, strResult1, strResult2)
	printOutputOperatorLogikaOR(result1, result2, strResult1, strResult2)
	fmt.Println("==============================================")

	fmt.Println()
	fmt.Println("Operator Relasi")
	result1 = x >= y
	result2 = x <= y
	if result1 {
		fmt.Println("Hasil dari", x, ">=", y, "--> 1")
		strResult1 = 1
	} else {
		fmt.Println("Hasil dari", x, ">=", y, "--> 0")
		strResult1 = 0
	}
	if result2 {
		fmt.Println("Hasil dari", x, "<=", y, "--> 1")
		strResult2 = 1
	} else {
		fmt.Println("Hasil dari", x, "<=", y, "--> 0")
		strResult2 = 0
	}
	fmt.Println()

	fmt.Println("Operator Logika")
	printOutputOperatorLogikaAND(result1, result2, strResult1, strResult2)
	printOutputOperatorLogikaOR(result1, result2, strResult1, strResult2)
	fmt.Println("==============================================")
}

func printOutputOperatorLogikaAND(value1 bool, value2 bool, strValue1 int, strValue2 int) {
	if value1 && value2 {
		fmt.Println("Hasil dari", strValue1, "&&", strValue2, "--> 1")
	} else {
		fmt.Println("Hasil dari", strValue1, "&&", strValue2, "--> 0")
	}
}

func printOutputOperatorLogikaOR(value1 bool, value2 bool, strValue1 int, strValue2 int) {
	if value1 || value2 {
		fmt.Println("Hasil dari", strValue1, "||", strValue2, "--> 1")
	} else {
		fmt.Println("Hasil dari", strValue1, "||", strValue2, "--> 0")
	}
}
