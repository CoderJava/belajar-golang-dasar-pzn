package main

import "belajar-golang-dasar/helper"


func main() {
	helper.SayHello("Yudi Setiawan")
}

// Ini GO111MODULE-nya dalam kondisi off agar bisa berjalan.
// Untuk menghidupkan GO111MODULE-nya gunakan command berikut.
// go env -w GO111MODULE=on
//
// Untuk mematikan-nya gunakan perintah berikut.
// go env -w GO111MODULE=off
//
// Untuk melihat status env-nya gunakan perintah berikut.
// go env
