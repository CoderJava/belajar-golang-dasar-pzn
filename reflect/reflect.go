package main

import (
	"fmt"
	"reflect"
)

func main() {
	sample := Sample{"Eko", 28}		
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(1).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min")) // empty string jika tag-nya tidak ada.

	fmt.Println(isValid(sample)) // true
	sample.Name = ""
	fmt.Println(isValid(sample)) // false
	fmt.Println()

	contoh := ContohLagi{"", ""} // karena didalam field struct-nya tidak ada StructTag
	fmt.Println(isValid(contoh)) // true
}

type Sample struct {
	Name string `required:"true" max:"10"`
	Age int
}

func isValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name string
	Description string
}