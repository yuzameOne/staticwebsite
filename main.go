package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Server start ")

	staticfile := http.FileServer(http.Dir("static"))

	http.Handle("/", staticfile)

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Index Page")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "About Page")
	})

	http.ListenAndServe(":8080", nil)
}
