package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func contador(endpoint string) {
	flag := os.Getenv("COUNTER_HIT_GOLANG")
	count, err := strconv.Atoi(flag)

	if err != nil {
		panic(err)
	}

	newValue := strconv.Itoa(count + 1)
	os.Setenv("COUNTER_HIT_GOLANG", newValue)
	fmt.Println("Endpoint Hit:", endpoint, newValue)
}

func resp_health(w http.ResponseWriter, r *http.Request) {
	mapD := map[string]string{"response": "ok"}
	mapB, _ := json.Marshal(mapD)

	fmt.Fprintf(w, string(mapB))
	contador("health")
}

func resp_unhealthy(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	mapD := map[string]string{"response": "cof...cof"}
	mapB, _ := json.Marshal(mapD)

	fmt.Fprintf(w, string(mapB))
	contador("unhealthy")
}

func resp_pong(w http.ResponseWriter, r *http.Request) {
	mapD := map[string]string{"ping": "pong"}
	mapB, _ := json.Marshal(mapD)

	fmt.Fprintf(w, string(mapB))
	contador("ping-pong")
}

func handleRequests() {
	http.HandleFunc("/health", resp_health)
	http.HandleFunc("/unhealthy", resp_unhealthy)
	http.HandleFunc("/ping", resp_pong)
	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {
	handleRequests()
}
