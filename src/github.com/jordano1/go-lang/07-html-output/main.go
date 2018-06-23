package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Jordan"
	d := []string{name}
	title := append(d, "'s site")
	s := strings.Join(title, "")
	fmt.Println("Title:", title)
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>` + s + `</title>
</head>
<body>
    <p>` + name + `</p>
</body>
</html>`

	fmt.Println(html)
}
