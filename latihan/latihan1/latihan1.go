package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println()
	fmt.Println("       BIODATA PRIBADY SAYA        ")
	fmt.Println("===================================")
	fmt.Print("Nama Lengkap   :")
	scanner.Scan()
	namaLengkap := scanner.Text()

	fmt.Print("Kota Kelahiran :")
	scanner.Scan()
	kotaKelahiran := scanner.Text()

	fmt.Print("Tanggal Lahir  :")
	scanner.Scan()
	tanggalLahir := scanner.Text()

	fmt.Print("Alamat         :")
	scanner.Scan()
	alamat := scanner.Text()

	fmt.Print("No.Telp        :")
	scanner.Scan()
	noTelp := scanner.Text()

	fmt.Println()
	output := `Perkenalkan nama saya adalah %s.
Saya lahir di %s.
Saya lahir pada tanggal %s.
Alamat saya %s.
Teman-teman boleh menghubungi saya di nomor telpon ini --> %s

Salam Persahabatan...
	`
	fmt.Printf(output, namaLengkap, kotaKelahiran, tanggalLahir, alamat, noTelp)
}
