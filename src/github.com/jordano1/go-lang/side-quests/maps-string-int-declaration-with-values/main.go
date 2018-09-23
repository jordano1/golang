package main

import "fmt"

var m map[string]int

func main() {

	m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	fmt.Println(m["one"] + m["five"])

}
