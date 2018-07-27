package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

type course struct {
	Number, Name, Units string
}
type semester struct {
	Term    string
	Courses []course
}
type year struct {
	Fall, Spring, Summer semester
}

func main() {
	y := year{
		Fall: semester{
			Term: "fall",
			Courses: []course{
				course{"school is for fools", "school is for fools", "school is for fools"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"school is for fools", "school is for fools", "school is for fools"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}
