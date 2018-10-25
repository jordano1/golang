package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggo")

}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cato")
}

func e(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "root")
}

func main() {

	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)
	http.HandleFunc("/", e)

	http.ListenAndServe(":8080", nil)
}
