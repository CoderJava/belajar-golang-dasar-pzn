package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	daftarNamaPaket := map[int8]string{
		1: "PAKET HEMAT",
		2: "PAKET NASI",
		3: "PAKET SPESIAL",
	}
	daftarHargaPaket := map[int8]int16{
		1: 7500,
		2: 10000,
		3: 15000,
	}

	fmt.Println()
	fmt.Print("Masukkan Kode [1..3]?\t: ")
	scanner.Scan()
	strKode := scanner.Text()
	kode, error := strconv.Atoi(strKode)
	if error != nil {
		fmt.Println("Inputan kode tidak valid. Harus bernilai angka ya.")
		return
	} else if kode < 1 || kode > 3 {
		fmt.Println("Inputan kode tidak valid. Harus angka dari 1 s.d. 3.")
		return
	}
	namaPaketYangDipilih := daftarNamaPaket[int8(kode)]
	hargaPaketYangDipilih := daftarHargaPaket[int8(kode)]

	fmt.Print("Jumlah Beli\t\t: ")
	scanner.Scan()
	strJumlahBeli := scanner.Text()
	jumlahBeli, error := strconv.Atoi(strJumlahBeli)
	if error != nil {
		fmt.Println("Inputan jumlah beli tidak valid. Harus angka ya.")
		return
	}
	total := jumlahBeli * int(hargaPaketYangDipilih)
	ppn := int(float32(total) * 0.10)
	jumlahBayar := total + ppn

	fmt.Print("Kode Kasir\t\t: ")
	scanner.Scan()
	strKodeKasir := scanner.Text()

	fmt.Print("Nama Kasir\t\t: ")
	scanner.Scan()
	namaKasir := scanner.Text()

	fmt.Println()
	fmt.Println("                 SEJAHTERA CAFE              ")
	fmt.Println("                Jl. Juang No.168             ")
	fmt.Println("              Telp.7236574-7236574           ")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("\t" + namaKasir + " " + strKodeKasir)
	fmt.Println("\t" + namaPaketYangDipilih)
	fmt.Println("\t"+strJumlahBeli, "X Rp.", hargaPaketYangDipilih)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Printf("\tTotal        : Rp. %d\n", total)
	fmt.Printf("\tPPN 10%%      : Rp. %d\n", ppn)
	fmt.Printf("\tJumlah Bayar : Rp. %d\n", jumlahBayar)

	fmt.Print("\tBayar        : Rp. ")
	scanner.Scan()
	strBayar := scanner.Text()
	bayar, error := strconv.Atoi(strBayar)
	if error != nil {
		fmt.Println("Inputan jumlah bayar tidak valid.")
		return
	} else if bayar < jumlahBayar {
		fmt.Println("Uang pembayaran tidak mencukupi.")
		return
	}
	uangKembali := bayar - jumlahBayar
	fmt.Printf("\tKembali      : Rp. %d\n", uangKembali)
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("            SELAMAT MENIKMATI                ")
	fmt.Println("              TERIMA KASIH                   ")
	fmt.Println()
}
