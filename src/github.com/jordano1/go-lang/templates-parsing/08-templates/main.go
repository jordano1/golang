package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {

	err := tpl.Execute(os.Stdout, "templates/*")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("line 15: ")
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nline 20: ", tpl.Execute(os.Stdout, nil))
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nline 24: ")
	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("\nline 29: ")
	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
