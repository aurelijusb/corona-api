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
	mux.Get("/api/v1/dates", dates)
	mux.Get("/api/v1/historical", historical)
	mux.Get("/api/v1/latest", latest)
	mux.Get("/", usage)
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

func usage(w http.ResponseWriter, req *http.Request) {
	app.RespondJSON(map[string]interface{}{
		"endpoints": []string{
			"/ping",
			"/api/v1/raw",
			"/api/v1/raw/{fileName}",
			"/api/v1/dates",
			"/api/v1/historical",
			"/api/v1/latest",
		},
		"source": map[string]string{
			"data": "https://sam.lrv.lt/koronavirusas",
			"code": "https://github.com/aurelijusb/corona-api",
		},
		"version": "0.1.0",
	}, w)
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

func dates(w http.ResponseWriter, req *http.Request) {
	files, err := app.GetFiles(app.GetEnv("DATA_PATH", "data/"))
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	var times []string
	for _, fileName := range files {
		times = append(times, app.FileNameToDate(fileName))
	}
	app.RespondJSON(times, w)
}

func historical(w http.ResponseWriter, req *http.Request) {
	dataPath := app.GetEnv("DATA_PATH", "data/")
	files, err := app.GetFiles(dataPath)
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	statistics := []app.CoronaReport{}
	for _, fileName := range files {
		content, err := app.ReadFile(dataPath, fileName)
		if err != nil {
			app.RespondWithError(err, w, req)
			return
		}
		statistic := app.ExtractData(string(content), fileName)
		statistics = append(statistics, statistic)
	}

	app.RespondJSON(statistics, w)
}

func latest(w http.ResponseWriter, req *http.Request) {
	dataPath := app.GetEnv("DATA_PATH", "data/")
	files, err := app.GetFiles(dataPath)
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	fileName := files[len(files)-1] // Latest
	content, err := app.ReadFile(dataPath, fileName)
	if err != nil {
		app.RespondWithError(err, w, req)
		return
	}
	statistic := app.ExtractData(string(content), fileName)

	app.RespondJSON(statistic, w)
}
