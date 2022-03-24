package main

import (
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
