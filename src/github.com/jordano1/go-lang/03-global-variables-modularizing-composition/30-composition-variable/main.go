package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

type course struct {
	Number, Name, Credits string
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
			Term: "Fall",
			Courses: []course{
				course{"1", "Dumputer Dcience", "3"},
				course{"2", "Bromputer Brience", "3"},
				course{"3", "Computer Science", "3"},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				course{"3", "Computer Science", "3"},
				course{"2", "Bromputer Brience", "3"},
				course{"1", "Dumputer Dcience", "3"},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				course{"3", "Computer Science", "3"},
				course{"1", "Dumputer Dcience", "3"},
				course{"2", "Bromputer Brience", "3"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", y)
	if err != nil {
		log.Fatalln(err)
	}
}
