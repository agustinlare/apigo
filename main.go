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
	http.HandleFunc("/checklist", checklist)
	http.HandleFunc("/checker", checker)

	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {
	handleRequests()
}
