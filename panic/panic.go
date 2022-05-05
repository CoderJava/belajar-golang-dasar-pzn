package main

import "fmt"

func main() {
	// runApp(true)
	runApp2(true)
	fmt.Println("Yudi")
}

func endApp() {
	// recover() berfungsi agar program tidak berhenti ketika panic. Dan menangkap pesan error dari panic.
	// recover() harus diletakkan didalam function defer.
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(isError bool) {
	defer endApp()	
	if isError {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func runApp2(isError bool) {
	defer endApp()
	if isError {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}
