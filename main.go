package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/unhealthy", unhealthy)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/front", front)
	http.HandleFunc("/checker", checker)
	http.HandleFunc("/password", rickroll)
	http.HandleFunc("/upload", upload)

	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {
	// listECR()
	initEnvs()
	handleRequests()
}
