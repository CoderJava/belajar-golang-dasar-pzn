package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const hargaBerasPerKilo = 4500

	fmt.Println()
	fmt.Println("                Toko Beras Murah               ")
	fmt.Println("===============================================")
	fmt.Print("Jumlah beli beras (kg)\t: ")
	scanner.Scan()
	strJumlahBeliBeras := scanner.Text()

	// Hitung total belanja
	jumlahBeliBeras, error := strconv.Atoi(strJumlahBeliBeras)
	if error != nil {
		fmt.Println("Nilai jumlah beli beras tidak valid. Harus angka ya.")
		return
	} else if jumlahBeliBeras < 1 {
		fmt.Println("Maaf Anda tidak dapat bonus")
		return
	}
	totalBelanja := hargaBerasPerKilo * jumlahBeliBeras

	fmt.Println("Total belanja Anda\t: Rp.", totalBelanja)

	if jumlahBeliBeras >= 23 {
		fmt.Println("Bonus 1 liter ice cream chocolatte")
	}

	fmt.Println("===============================================")
	fmt.Println("         Terima kasih atas kunjungannya        ")
}
