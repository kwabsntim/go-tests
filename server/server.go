package main

import (
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello this is the homepage"))
}
func AboutPage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello this is the about page"))
}

func main() {

	http.HandleFunc("/", HomePage)
	http.HandleFunc("/about", AboutPage)
	http.ListenAndServe(":8080", nil)
}
