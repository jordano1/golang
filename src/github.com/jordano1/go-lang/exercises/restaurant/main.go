package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name, Desc string
	Price      float64
}
type food struct {
	Hour string
	Item []item
}

type menu []food

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	m := menu{
		food{
			Hour: "Breakfast",
			Item: []item{
				item{
					Name:  "Eggs n bacon",
					Desc:  "Has Eggs and bacon",
					Price: 10.99,
				},
				item{
					Name:  "Bacon Coffee",
					Desc:  "Has Coffee and some strips of bacon cuz you're a weird hipster on a weird hipster diet",
					Price: 20.99,
				},
				item{
					Name:  "Meat Mountain",
					Desc:  "You ordered all our meat!",
					Price: 45.99,
				},
			},
		},
		food{
			Hour: "Lunch",
			Item: []item{
				item{
					Name:  "Grilled cheese",
					Desc:  "self explanatory",
					Price: 45.99,
				},
			},
		},
		food{
			Hour: "Dinner",
			Item: []item{
				item{
					Name:  "Spaghetti",
					Desc:  "everyone's second favorite",
					Price: 25.99,
				},
				item{
					Name:  "Pizza",
					Desc:  "everyone's favorite",
					Price: 15.99,
				},
				item{
					Name:  "Fetuccini",
					Desc:  "sassy sauce, sexy noodles",
					Price: 12.99,
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
