package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")

	flag := os.Getenv("COUNTER_HIT_GOLANG")
	count, err := strconv.Atoi(flag)

	if err != nil {
		panic(err)
	}

	newValue := strconv.Itoa(count + 1)
	os.Setenv("COUNTER_HIT_GOLANG", newValue)
	fmt.Println("Endpoint Hit: homePage", newValue)

	// if strings.Contains(newValue, "100") {
	// 	time.Sleep(5 * time.Second)
	// 	fmt.Println("Endpoint Hit by 5: homePage", newValue)
	// } else if strings.Contains(newValue, "40") {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Endpoint Hit by 2: homePage", newValue)
	// } else {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("Endpoint Hit by 0: homePage", newValue)
	// }

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("Server Up")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func main() {

	handleRequests()
}
