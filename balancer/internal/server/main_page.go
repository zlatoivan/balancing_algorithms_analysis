package server

import (
	"html/template"
	"log"
	"net/http"
)

type ViewData struct {
	Title   string
	Message string
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/template/index.html")
	err := t.Execute(w, "")
	if err != nil {
		log.Printf("t.Execute: %v", err)
	}
}
