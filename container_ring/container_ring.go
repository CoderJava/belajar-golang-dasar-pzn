package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	data := ring.New(5)

	// data.Value = "Eko"

	// var data2 = data.Next()
	// data2.Value = "Kurniawan"

	for i := 0; i < data.Len(); i++ {
		data.Value = "Data " + strconv.Itoa(i)
		data = data.Next()
	}

	// fmt.Println(*data)

	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
