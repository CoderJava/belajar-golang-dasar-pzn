package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 = int64(nilai32)
	var nilai8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Yudi"
	var char = name[0]
	var strChar = string(char)

	fmt.Println(name)
	fmt.Println(char)
	fmt.Println(strChar)
}