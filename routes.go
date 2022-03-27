package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	if getHealth() {
		responder(w, r, true, "true")
	} else {
		responder(w, r, false, "false")
	}
}

func unhealthy(w http.ResponseWriter, r *http.Request) {
	responder(w, r, false, "cof...cof")
}

func ping(w http.ResponseWriter, r *http.Request) {
	responder(w, r, true, "pong")
}

func rickroll(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.youtube.com/watch?v=dQw4w9WgXcQ", 302)
	responder(w, r, true, "Never gonna give you up!")
}

func checker(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(err.Error())
	}

	if len(r.Form["ipcheck"][0]) > 0 {
		reply, err := connReach(r.Form["ipcheck"][0])
		responder(w, r, true, reply)

		if err != nil {
			log.Println(err)
		}

	} else if len(r.Form["mongodb"][0]) > 0 {

		reply, err := mongodb(r.Form["mongodb"][0])
		responder(w, r, true, reply)

		if err != nil {
			log.Println(err)
		}

	} else if len(r.Form["dns"][0]) > 0 {

		reply, err := dnsResolver(r.Form["dns"][0])
		responder(w, r, true, reply)

		if err != nil {
			log.Println(err)
		}

	} else {
		responder(w, r, true, "Empty request")
	}
}
