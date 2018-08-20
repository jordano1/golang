package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Name, Desc, Meal string
	Price            float64
}

type menus []menu

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	m := menus{
		menu{
			Name:  "Oatmeal",
			Desc:  "yum yum",
			Meal:  "Breakfast",
			Price: 4.95,
		},
		menu{
			Name:  "Hamburger",
			Desc:  "Delicous good eating for you",
			Meal:  "Lunch",
			Price: 6.95,
		},
		menu{
			Name:  "Pasta Bolognese",
			Desc:  "From Italy delicious eating",
			Meal:  "Dinner",
			Price: 7.95,
		},
	}
	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
