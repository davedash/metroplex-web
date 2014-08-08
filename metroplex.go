package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

type Node struct {
	Hostname string
	Services []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stderr, http.DefaultServeMux))
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}
