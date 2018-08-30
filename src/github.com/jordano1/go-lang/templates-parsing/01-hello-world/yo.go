package main

import (
	"fmt"
	"sort"
)

func main() {

	d := []string{"c", "b", "f", "a"}
	fmt.Println(d)
	sort.Strings(d)
	fmt.Println(d)

}
