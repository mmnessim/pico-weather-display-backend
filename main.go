package main

import (
	"fmt"
	"log"
	"net/http"
	"server/weather"
	"time"

	"github.com/mmnessim/go-env"
)

var e = env.New(".env")
var apiKey = e.Get("API_KEY")

const (
	LAT  = 35.9956
	LONG = -78.9002
)

func SendWeather(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Printf("Recieved request at %s\n", now.String())
	l := weather.GetLatAndLong("27713")
	weather := weather.GetWeatherWithLatAndLong(l.Lat, l.Long)
	fmt.Fprintf(w, "cur=%s high=%s low=%s weather=%s precip=%s",
		weather.Current, weather.High, weather.Low, weather.Weather, weather.Percipitation)
}

func SendWeatherWithZip(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Printf("Recieved request at %s\n", now.String())
	zip := r.URL.Query().Get("zip")
	l := weather.GetLatAndLong(zip)
	weath := weather.GetWeatherWithLatAndLong(l.Lat, l.Long)
	fmt.Fprintf(w, "cur=%s high=%s low=%s weather=%s precip=%s",
		weath.Current, weath.High, weath.Low, weath.Weather, weath.Percipitation)

}

func main() {
	_ = weather.GetLatAndLong("27713")

	weather.GetWeather()
	http.HandleFunc("/", SendWeather)
	//http.HandleFunc("/lat", GetLatAndLong)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
