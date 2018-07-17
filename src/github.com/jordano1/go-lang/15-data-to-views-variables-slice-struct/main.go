package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type sage struct {
	Name   string
	Saying string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	jesus := sage{
		Name:   "jesus",
		Saying: "hypocrasy is my game",
	}
	buddha := sage{
		Name:   "buddha",
		Saying: "hella",
	}
	sages := []sage{jesus, buddha}

	err := tpl.Execute(os.Stdout, sages)
	if err != nil {
		log.Fatalln(err)
	}
}
