package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	host := getEnv("SERVER_HOST", "0.0.0.0")
	port := getEnv("SERVER_PORT", "80")

	fmt.Printf("Listening on %s:%s\n", host, port)
	http.HandleFunc("/ping", ping)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
