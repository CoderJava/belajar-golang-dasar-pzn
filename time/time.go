package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println()

	utc := time.Date(2022, 5, 11, 6, 49, 10, 20, time.UTC)
	fmt.Println(utc)
	fmt.Println()

	// January / Jan / january / jan / 01 / -1 (etc) are for month
	// 02 / _2 are for day of month
	// 15 / 03 / _3 / PM / P / pm /p are for hour & meridian (3pm)
	// 04 / _4 are for minutes
	// 05 / _5 are for seconds
	// 2006 / 06 are for year
	// -0700 / 07:00 / MST are for timezone
	// 999999999 / .000000000 etc are for partial seconds (I think the distinction is if trailing zeros are removed)
	// Mon / Monday are day of the week (which 01-02-2006 actually was)
	layout := "02-01-2006 03:04:05" // time.RFC3339
	parse, err := time.Parse(layout, "11-05-2022 07:08:30")
	if err == nil {
		fmt.Println(parse)
	} else {
		fmt.Println(err.Error())
	}
}
