package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type items struct {
	Wisdom []sage
	Car    []car
}

type sage struct {
	Name   string
	Saying string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	j := sage{
		Name:   "jesus",
		Saying: "hypocrasy is my game",
	}
	b := sage{
		Name:   "buddha",
		Saying: "hella",
	}
	f := car{
		Manufacturer: "Ford",
		Model:        "f150",
		Doors:        2,
	}
	d := car{
		Manufacturer: "farts",
		Model:        "f2000",
		Doors:        4,
	}
	k := car{
		Manufacturer: "fefefe",
		Model:        "asdfasdfdasf",
		Doors:        67,
	}
	sages := []sage{j, b}
	cars := []car{f, d, k}
	data := items{
		Wisdom: sages,
		Car:    cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
