package main

import (
	"html/template"
	"fmt"
	"net/http"
)

type Node struct {
	Hostname string
	Services []string
}

func handler(w http.ResponseWriter, r *http.Request) {
	if (r.URL.Path != "/") {
	  errorHandler(w, r, http.StatusNotFound)
		return
	}
	t, _ := template.ParseFiles("templates/home.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}
