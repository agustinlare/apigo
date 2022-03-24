package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type response struct {
	Endpoint string `json:"endpoint"`
	Ip       string `json:"ip"`
	Counter  int    `json:"counter"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}

func get_ip(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")

	if forwarded != "" {
		return forwarded
	}

	return r.RemoteAddr
}

func hit_counter(endpoint string) int {
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

func responder(w http.ResponseWriter, r *http.Request, b bool, s string) {
	w.Header().Add("Content-Type", "application/json")

	status := 200

	if !b {
		w.WriteHeader(http.StatusInternalServerError)
		status = 500
	}

	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')

	if lastSlash < 0 {
		lastSlash = 0
	}

	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash
	calleIp := get_ip(r)
	countHp := hit_counter(funcName[lastDot+1:])

	mapD := response{
		Endpoint: funcName[lastDot+1:],
		Ip:       calleIp,
		Counter:  countHp,
		Status:   status,
		Message:  s,
	}

	mapB, _ := json.Marshal(mapD)

	w.Write(mapB)
	fmt.Println(string(mapB))
}

// func PrettyString(str string) (string, error) {
// 	var prettyJSON bytes.Buffer
// 	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
// 			return "", err
// 	}
// 	return prettyJSON.String(), nil
// }
