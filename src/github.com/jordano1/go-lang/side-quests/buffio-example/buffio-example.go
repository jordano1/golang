package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := `Gonna poop my pants`
	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
