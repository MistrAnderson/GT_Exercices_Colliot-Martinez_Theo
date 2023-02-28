package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port = ":8080"

func main() {

	fmt.Println("(http://localhost:8080)", "Server started on port ", port)

	// Gère la route "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	})

	// Gère la route "/un"
	http.HandleFunc("/un", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/un" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/un.html"))
		tmpl.Execute(w, nil)
	})

	// Gère la route "/deux"
	http.HandleFunc("/deux", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/deux" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}

		tmpl := template.Must(template.ParseFiles("templates/deux.html"))
		tmpl.Execute(w, nil)
	})

	// Gère la route "/trois"
	http.HandleFunc("/trois", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/trois" {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		tmpl := template.Must(template.ParseFiles("templates/trois.html"))
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8080", nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl := template.Must(template.ParseFiles("templates/error.html"))
		tmpl.Execute(w, nil)
	}
}
