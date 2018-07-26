package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type user struct {
	Name  string
	Motto string
	Admin bool
}

func main() {
	u1 := user{
		Name:  "Jordan",
		Motto: "I poop first, and ask questions later",
		Admin: false,
	}
	u2 := user{
		Name:  "brosef",
		Motto: "pizza hut is my favorite place to kiss cute dudes",
		Admin: true,
	}
	u3 := user{
		Name:  "",
		Motto: "NOICE!!!",
		Admin: true,
	}
	u4 := user{
		Name:  "tiberius",
		Motto: "I farted",
		Admin: true,
	}
	//i like this format a lot
	users := []user{u1, u2, u3, u4}
	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
