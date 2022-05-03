package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const tarifTetap = 200000
	const tarifJamBerikutnya = 0.25

	fmt.Println()
	fmt.Println("              Studio Al Izzah             ")
	fmt.Println("==========================================")
	fmt.Println("Isi Daftar Penyewa Berikut")
	fmt.Print("Nama Group        :")
	scanner.Scan()
	scanner.Text()

	fmt.Print("Alamat            :")
	scanner.Scan()
	scanner.Text()

	fmt.Print("No.Telp           :")
	scanner.Scan()
	scanner.Text()

	fmt.Println("==========================================")
	fmt.Print("Lama Rental (jam) :")
	scanner.Scan()
	strLamaRental := scanner.Text()

	// Hitung nilai total bayar
	lamaRental, error := strconv.Atoi(strLamaRental)
	if error != nil {
		fmt.Println("Nilai lama rental tidak valid. Harus bernilai angka ya.")
		return
	} else if lamaRental < 1 {
		fmt.Println("Nilai lama rental minimal harus 1 jam ya.")
		return
	}

	var totalBayar int
	if lamaRental == 1 {
		totalBayar = tarifTetap
	} else {
		lamaRentalTambahan := lamaRental - 1
		biayaRentalTambahan := int(float64(lamaRentalTambahan*tarifTetap) * tarifJamBerikutnya)
		totalBayar = tarifTetap + biayaRentalTambahan
	}
	fmt.Println("Total Bayar       :Rp.", totalBayar)
	fmt.Println("==========================================")
}
