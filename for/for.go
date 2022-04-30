package main

import "fmt"

func main() {

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	// names := []string{"Eko", "Kurniawan", "Khannedy"}

	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, name := range names {
	// 	fmt.Println("index:", index, " & name:", name)
	// }

	// for _, name := range names {
	// 	fmt.Println("name:", name)
	// }

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
