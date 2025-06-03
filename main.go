package main

import (
	"fmt"
	"log"
	"net/http"
	"server/weather"

	"github.com/mmnessim/go-env"
)

var e = env.New(".env")
var apiKey = e.Get("API_KEY")

const (
	LAT  = 35.9956
	LONG = -78.9002
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Request Info:\n")
	fmt.Fprintf(w, "Go Method: %s\n", r.Method)
	fmt.Fprintf(w, "Go URL: %s\n", r.URL.String())
	fmt.Fprintf(w, "Go Header: %v\n", r.Header)
	fmt.Fprintf(w, "Go RemoteAddr: %s\n", r.RemoteAddr)
}

func SendWeather(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recieved request")
	weather := weather.GetWeather()
	fmt.Fprintf(w, "cur=%s high=%s low=%s weather=%s", weather.Current, weather.High, weather.Low, weather.Weather)
}

func main() {
	weather.GetWeather()
	http.HandleFunc("/", SendWeather)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
