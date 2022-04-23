package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool
	
	var noKtpYudi NoKTP = "1234567890"
	var marriedStatus Married = true
	fmt.Println(noKtpYudi)
	fmt.Println(marriedStatus)
}