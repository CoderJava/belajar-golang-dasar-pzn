package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println("        Toko Elektronik Sejahtera         ")
	fmt.Println("==========================================")
	fmt.Print("Nama Barang    :")
	scanner.Scan()

	fmt.Print("Harga Barang   :Rp. ")
	scanner.Scan()
	strHargaBarang := scanner.Text()
	hargaBarang, error := strconv.Atoi(strHargaBarang)
	if error != nil {
		fmt.Print("Nilai harga barang tidak valid. Harus angka ya.")
		return
	} else if hargaBarang <= 0 {
		fmt.Print("Nila harga barang harus diatas 0.")
		return
	}

	fmt.Print("Jumlah Beli    :")
	scanner.Scan()
	strJumlahBeli := scanner.Text()
	jumlahBeli, error := strconv.Atoi(strJumlahBeli)
	if error != nil {
		fmt.Print("Nilai jumlah beli tidak valid. Harus angka ya.")
		return
	} else if jumlahBeli < 0 {
		fmt.Print("Nilai jumlah beli harus diatas 0.")
		return
	}

	// Hitung jumlah bayar, PPN, dan total bayar
	jumlahBayar := hargaBarang * jumlahBeli
	ppn := int(float64(jumlahBayar) * 0.10)
	totalBayar := jumlahBayar + ppn

	fmt.Println("Jumlah Bayar   :Rp.", jumlahBayar)
	fmt.Println("PPN (10%)      :Rp.", ppn)
	fmt.Println("Total Bayar    :Rp.", totalBayar)
	fmt.Println("==========================================")

	fmt.Print("Uang Bayar     :Rp. ")
	scanner.Scan()
	strUangBayar := scanner.Text()
	uangBayar, error := strconv.Atoi(strUangBayar)
	if error != nil {
		fmt.Print("Nilai uang bayar tidak valid. Harus angka ya.")
		return
	} else if uangBayar < totalBayar {
		fmt.Print("Waduh... Uang kamu tidak cukup nih.")
		return
	}

	// Hitung uang kembali
	uangKembali := uangBayar - totalBayar
	fmt.Println("Uang Kembali   :Rp. ", uangKembali)
	fmt.Println("==========================================")
	fmt.Print("       TERIMA KASIH ATAS KUNJUNGANNYA       ")
}
