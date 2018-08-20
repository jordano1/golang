package main

import (
	"html/template"
	"log"
	"os"
)

//contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

type time struct {
	Breakfast, Lunch, Dinner string
}

type menu struct {
	Time []time
}

type restaurant struct {
	Name string
	Menu []menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	r := restaurant{
		Name: "Joes",
		Menu: []menu{
			menu{
				Time:time{
					time{}
				},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", r)
	if err != nil {
		log.Fatalln(err)
	}
}
