package main

import (
	"html/template"
	"net/http"
)

type Node struct {
	Hostname string
	Services []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	r.URL
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
