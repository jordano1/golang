package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

func main() {
	tpl, err := template.ParseFiles("../07-html-output/tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("error creating file: ", err)
	}
	fmt.Println(nf)

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
