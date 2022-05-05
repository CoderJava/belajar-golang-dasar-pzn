package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println("                W A R N E T               ")
	fmt.Println("==========================================")
	fmt.Print("Nama Pengunjung\t\t\t: ")
	scanner.Scan()

	fmt.Print("Durasi Sewa (Jam)\t\t: ")
	scanner.Scan()
	strDurasiSewa := scanner.Text()
	durasiSewa, error := strconv.Atoi(strDurasiSewa)
	if error != nil {
		fmt.Println("Nilai durasi sewa tidak valid. Minimal 1 jam ya.")
		return
	}

	fmt.Print("Keterangan\t\t\t: ")
	scanner.Scan()
	keterangan := strings.ToLower(scanner.Text())

	totalPembayaran := 0
	hargaDiskon := 0
	switch {
	case keterangan == "umum":
		{
			// Handle untuk pelanggan umum
			hargaSewaPerJam := 5000
			diskonSewaPerJam := 0.125
			totalPembayaran = durasiSewa * hargaSewaPerJam
			hargaDiskon = int(float64(totalPembayaran) * diskonSewaPerJam)
			totalPembayaran += hargaDiskon
		}
	case keterangan == "p":
		{
			// Handle untuk pelanggan dengan kode "P"
			hargaSewaPerJam := 4000
			var diskonSewaPerJam float32
			if durasiSewa >= 5 {
				diskonSewaPerJam = 0.5
			} else if durasiSewa >= 3 {
				diskonSewaPerJam = 0.3
			} else {
				diskonSewaPerJam = 0.125
			}
			totalPembayaran = durasiSewa * hargaSewaPerJam
			hargaDiskon = int(float64(totalPembayaran) * float64(diskonSewaPerJam))
			totalPembayaran += hargaDiskon
		}
	default:
		{
			fmt.Println("Keterangan tidak valid.")
			return
		}
	}
	fmt.Println("Diskon Yang Diperoleh\t\t: Rp.", hargaDiskon)
	fmt.Println("Total Pembayaran\t\t: Rp.", totalPembayaran)
	fmt.Println("==========================================")
	fmt.Print("Uang Bayar\t\t\t: Rp. ")
	scanner.Scan()
	strUangBayar := scanner.Text()

	uangBayar, error := strconv.Atoi(strUangBayar)
	if error != nil {
		fmt.Println("Nilai uang bayar tidak valid. Harus angka ya.")
		return
	} else if uangBayar < totalPembayaran {
		fmt.Println("Uang bayar tidak cukup.")
		return
	}
	uangKembali := uangBayar - totalPembayaran
	fmt.Println("Uang Kembali\t\t\t: Rp.", uangKembali)
	fmt.Println("==========================================")
	fmt.Println("     TERIMA KASIH ATAS KUNJUNGAN ANDA     ")
}
