package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	// go run os.go Yudi Setiawan
	// fmt.Println(args[1])
	// fmt.Println(args[2])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}
