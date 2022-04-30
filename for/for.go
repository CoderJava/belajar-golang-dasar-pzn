package main

import "fmt"

func main() {

	// Loop cara 1
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	// Loop cara 2
	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke", counter)
	// }

	// names := []string{"Eko", "Kurniawan", "Khannedy"}

	// Iteration collection cara 1
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// Iteration collection cara 2
	// for index, name := range names {
	// 	fmt.Println("index:", index, " & name:", name)
	// }

	// Iteration collection cara 3
	// for _, name := range names {
	// 	fmt.Println("name:", name)
	// }

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	// Iteration map
	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
