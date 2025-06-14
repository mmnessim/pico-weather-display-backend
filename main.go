package main

import (
	"fmt"
	"log"
	"net/http"
	"server/weather"
	"time"
)

const (
	LAT  = 35.9112
	LONG = -78.9178
)

func SendWeather(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	fmt.Printf("Recieved request at %s\n", now.String())
	//fmt.Printf("URL: %s\n", r.URL)
	zipcode := r.URL.Query().Get("zipcode")
	l := weather.GetLatAndLong(zipcode)
	weath := weather.GetWeatherWithLatAndLong(l.Lat, l.Long)
	fmt.Println(weath)
	fmt.Fprintf(w, "cur=%s high=%s low=%s weather=%s precip=%s",
		weath.Current, weath.High, weath.Low, weath.Weather, weath.Percipitation)
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
	//_ = weather.GetLatAndLong("27713")

	weather.GetWeather()
	http.HandleFunc("/", SendWeather)
	//http.HandleFunc("/lat", GetLatAndLong)
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
