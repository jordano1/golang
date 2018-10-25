package main

import (
	"io"
	"net/http"
)

type hotdog int

func (j hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo")

}

type hotcat int

func (j hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cato")
}

type hote int

func (j hote) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cato")
}
func main() {
	var d hotdog
	var c hotcat
	var e hote

	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat/", c)
	mux.Handle("/", e)

	http.ListenAndServe(":8080", mux)
}
