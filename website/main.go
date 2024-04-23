package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func home(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:   "index page",
		Content: "welcome to homepage",
	}
	render(w, "index", data)
}

func render(w http.ResponseWriter, page string, data interface{}) {
	t, err := template.ParseFiles(fmt.Sprintf("./templates/%s.html", page))
	if err != nil {
		panic(err)
	}
	t.Execute(w, data)
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))

	http.HandleFunc("/", home)

	http.ListenAndServe(":3090", nil)
}
