package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the root url!")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the dog url")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the me url, I'm jordan!")
}
