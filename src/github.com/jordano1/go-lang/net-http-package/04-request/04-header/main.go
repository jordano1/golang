package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Jordan-key", "this is from Jordan")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(res, "<h1 style='color:blue;'>some code brosef</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
