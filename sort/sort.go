package main

import (
	"fmt"
	"sort"
)

func main() {
	friends := UserSlice{
		{"Teman Eko", 30},
		{"Teman Joko", 35},
		{"Teman Budi", 31},
		{"Teman Rudi", 25},
	}

	users := []User {
		{"Eko", 30},
		{"Joko", 35},
		{"Budi", 31},
		{"Rudi", 25},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users)) // konversi menjadi slice
	fmt.Println(users)
	fmt.Println()

	fmt.Println(friends)
	sort.Sort(friends)
	fmt.Println(friends)
}

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	// return value[i].Name < value[j].Name
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

// func (value UserSlice) Len() int           { return len(value) }
// func (value UserSlice) Swap(i, j int)      { value[i], value[j] = value[j], value[i] }
// func (value UserSlice) Less(i, j int) bool { return value[i] < value[j] }
