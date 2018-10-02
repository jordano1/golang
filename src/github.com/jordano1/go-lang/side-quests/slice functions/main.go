package main

import "fmt"

func main() {
	f := []int{2, 2, 2}
	sum(3)
	sub(7)
	square(4)
	xsum(f)
}

func sum(n int64) {
	n += n
	fmt.Println(n)
}

func sub(n int64) {
	n -= n
	fmt.Println(n)
}

func square(n int64) {
	n *= n
	fmt.Println(n)
}

//slices
func xsum(n []int) {
	for _, z := range n {
		var c int
		for _, n := range n {
			c += n
			//adds all elements in c
			fmt.Println(c)
			fmt.Println(n)
			fmt.Println(z)
		}

		fmt.Println(n)
	}
}

func xsub(n []int) {

	fmt.Println(n)
}

func xsquare(n []int) {

	fmt.Println(n)
}
