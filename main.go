package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}
func main() {

	fmt.Println("Server start ")

	r := mux.NewRouter()

	cssHandler := http.FileServer(http.Dir("./static/css/"))
	imagesHandler := http.FileServer(http.Dir("./static/img/"))
	jsHandler := http.FileServer(http.Dir("./static/js"))

	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/img/", http.StripPrefix("/img/", imagesHandler))
	http.Handle("/js/", http.StripPrefix("/js/", jsHandler))
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}
