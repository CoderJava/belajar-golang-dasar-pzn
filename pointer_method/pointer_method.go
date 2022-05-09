package main

import "fmt"

func main() {
	eko := Man{Name: "Eko"}
	eko.Married()
	fmt.Println(eko)

	fmt.Println()
	eko.Married2()
	fmt.Println(eko)
}

type Man struct {
	Name string
}

// Nilai hanya berubah didalam method
func (man Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method", man.Name)
}

// Nilai-nya berubah karena menggunakan pointer
func (man *Man) Married2() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di Method2", man.Name)
}
