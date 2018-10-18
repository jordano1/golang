package main

import (
	"io"
	"net/http"
)

type hotdog int

func (j hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy")
	case "/cat":
		io.WriteString(res, "kitty")
	case "/":
		io.WriteString(res, "anus")

	}

}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
