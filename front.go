package main

import (
	"log"
	"net/http"
	"text/template"
)

func outputHTML(w http.ResponseWriter, filename string, data interface{}) {
	t, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := t.Execute(w, data); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func front(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	if r.Form["healthswitch"] != nil {
		responder(w, r, true, switchHealth())
	} else {
		responder(w, r, true, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	}

	myvar := map[string]interface{}{"HealthStatus": getHealth()}
	outputHTML(w, "static/index.html", myvar)
}
