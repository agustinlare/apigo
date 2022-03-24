package main

import (
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	responder(w, r, true, "ok")
}

func unhealthy(w http.ResponseWriter, r *http.Request) {
	responder(w, r, false, "cof...cof")
}

func ping(w http.ResponseWriter, r *http.Request) {
	responder(w, r, true, "pong")
}

func checklist(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
	responder(w, r, true, "checklist")
}

func checker(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

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
