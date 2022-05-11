package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("e([a-z])o")
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))
	fmt.Println()

	// -1 untuk max unlimited
	search := regex.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}

// https://github.com/google/re2/wiki/Syntax