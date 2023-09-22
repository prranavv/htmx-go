package main

import (
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/go-chi/chi/v5"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	r := chi.NewRouter()
	r.Get("/", newpage)
	r.Post("/add-film/", newfilm)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func newpage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	films := map[string][]Film{
		"Films": {
			{Title: "The GodFather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Rhidley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}
	tmpl.Execute(w, films)
}

func newfilm(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
}
