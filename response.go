package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type response struct {
	Endpoint string `json:"endpoint"`
	Ip       string `json:"ip"`
	Counter  int    `json:"counter"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}

func responder(w http.ResponseWriter, r *http.Request, b bool, s string) {
	status := 200

	if !b {
		w.WriteHeader(http.StatusInternalServerError)
		status = 500
	}

	callerUrl := r.URL.Path
	calleIp := getIp(r)
	countHp := hitCounter()

	mapD := response{
		Endpoint: callerUrl,
		Ip:       calleIp,
		Counter:  countHp,
		Status:   status,
		Message:  s,
	}

	mapB, _ := json.Marshal(mapD)

	log.Println(string(mapB))

	exceptions := []string{"front", "password"}

	if !stringInSlice(strings.TrimLeft(r.URL.Path, "/"), exceptions) {
		w.Header().Add("Content-Type", "application/json")
		w.Write(mapB)
	}
}
