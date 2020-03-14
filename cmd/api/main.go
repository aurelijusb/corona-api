package main

import (
	"fmt"
	"net/http"

	"github.com/aurelijusb/corona-api/internal/app"
	"github.com/gorilla/pat"
)

var mux *pat.Router = pat.New()

func init() {
	mux.Get("/ping", ping)
	mux.Get("/api/v1/raw/{file}", rawData)
	mux.Get("/api/v1/raw", rawList)
}

func main() {
	host := app.GetEnv("SERVER_HOST", "0.0.0.0")
	port := app.GetEnv("SERVER_PORT", "80")

	fmt.Printf("Listening on %s:%s\n", host, port)
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/raw", rawList)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), mux)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func rawList(w http.ResponseWriter, req *http.Request) {
	files, err := app.GetFiles(app.GetEnv("DATA_PATH", "data/"))
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	app.RespondJSON(files, w)
}

func rawData(w http.ResponseWriter, req *http.Request) {
	fileID := req.URL.Query().Get(":file")
	data, err := app.ReadFile(app.GetEnv("DATA_PATH", "data/"), fileID)
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	app.RespondHTML(data, w)
}
