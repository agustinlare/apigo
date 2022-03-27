package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func getIp(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}

func hitCounter() int {
	flag := os.Getenv("COUNTER_HIT_GOLANG")
	count, err := strconv.Atoi(flag)

	if err != nil {
		panic(err)
	}

	quick_math := count + 1
	newValue := strconv.Itoa(quick_math)
	os.Setenv("COUNTER_HIT_GOLANG", newValue)

	return quick_math
}

func getHealth() bool {
	boolValue, err := strconv.ParseBool(os.Getenv("HEALTCHECK_STATUS"))

	if err != nil {
		log.Println(err)
	}

	return boolValue
}

func switchHealth() string {
	os.Setenv("HEALTCHECK_STATUS", strconv.FormatBool(!getHealth()))

	return "Switched"
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
