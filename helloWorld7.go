package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World \n")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Earth \n")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/earth", handler2)
	http.ListenAndServe(":8080", nil)
}
