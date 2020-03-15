// Package app consists of helpers for HTTP server
package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// RespondWithError – output to the Browser with HTTP 500
func RespondWithError(err error, w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	corsAllowAll(w)
	noCache(w)
	text, jsonError := json.Marshal(map[string]string{"error": err.Error()})
	if jsonError != nil {
		fmt.Printf("Error: Unable to marshal for errro response: %s: %s", jsonError.Error(), err.Error())
	}
	fmt.Printf("Error: %s: %s", req.URL.String(), err.Error())
	fmt.Fprintf(w, "%s", text)
}

// RespondJSON – output to the Browser with JSON content type
func RespondJSON(data interface{}, w http.ResponseWriter) {
	text, jsonError := json.MarshalIndent(data, " ", " ")
	if jsonError != nil {
		fmt.Printf("Error: Unable to marshal response: %s: %#v", jsonError.Error(), data)
	}
	w.Header().Add("Content-type", "application/json")
	corsAllowAll(w)
	noCache(w)
	fmt.Fprintf(w, "%s", text)
}

// RespondHTML – return raw HTML
func RespondHTML(data []byte, w http.ResponseWriter) {
	w.Header().Add("Content-type", "text/html")
	corsAllowAll(w)
	cacheImmutable(w)
	fmt.Fprintf(w, "%s", data)
}

func corsAllowAll(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "*")
	w.Header().Add("Access-Control-Allow-Headers", "*")
}

func noCache(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "no-cache")
}

func cacheImmutable(w http.ResponseWriter) {
	w.Header().Add("Cache-Control", "public")
	w.Header().Add("Cache-Control", "immutable")
}

// GetEnv – os.Getenv with default value
func GetEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
