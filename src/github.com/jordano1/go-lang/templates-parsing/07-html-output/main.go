package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Jordan"
	d := []string{name}
	title := append(d, "'s site")
	s := strings.Join(title, "")
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
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error", err)
	}
	defer nf.Close()
	io.Copy(nf, strings.NewReader(html))

}
