package main

import (
	"fmt"
)

type person struct {
	f string
}

func (p *person) speak() {
	fmt.Println("hello")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		f: "Jordan",
	}
	saySomething(&p1)
	p1.speak()
}
