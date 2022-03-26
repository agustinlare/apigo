package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func health(w http.ResponseWriter, r *http.Request) {
	boolValue, err := strconv.ParseBool(os.Getenv("HEALTCHECK_STATUS"))
	if err != nil {
		log.Println(err)
	}

	if boolValue {
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

func front(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
	responder(w, r, true, "front")
}

func switcher(w http.ResponseWriter, r *http.Request) {
	reply, err := switchHealth()
	responder(w, r, true, reply)

	if err != nil {
		log.Println(err)
	}
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
