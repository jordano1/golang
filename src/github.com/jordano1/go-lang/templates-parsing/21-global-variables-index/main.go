package main

import (
	"html/template"
	"log"
	"os"
)

//declare tpl pointer
var tpl *template.Template

//initialize tpl pointer
func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	//slice of strings
	xs := []string{"one", "2", "3", "four", "five", "six"}
	//stdout
	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
