package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mmnessim/go-env"
)

var e = env.New(".env")
var apiKey = e.Get("API_KEY")

const (
	LAT  = 35.9956
	LONG = -78.9002
)

type Useful struct {
	Current       string
	High          string
	Low           string
	Weather       string
	Percipitation string
}

type Weather struct {
	Lat            int    `json:"lat"`
	Lon            int    `json:"lon"`
	Timezone       string `json:"timezone"`
	TimezoneOffset int    `json:"timezone_offset"`
	Current        struct {
		Dt         int     `json:"dt"`
		Sunrise    int     `json:"sunrise"`
		Sunset     int     `json:"sunset"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        float64 `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`
	Hourly []struct {
		Dt         int     `json:"dt"`
		Temp       float64 `json:"temp"`
		FeelsLike  float64 `json:"feels_like"`
		Pressure   int     `json:"pressure"`
		Humidity   int     `json:"humidity"`
		DewPoint   float64 `json:"dew_point"`
		Uvi        int     `json:"uvi"`
		Clouds     int     `json:"clouds"`
		Visibility int     `json:"visibility"`
		WindSpeed  float64 `json:"wind_speed"`
		WindDeg    int     `json:"wind_deg"`
		WindGust   float64 `json:"wind_gust"`
		Weather    []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Pop  int `json:"pop"`
		Rain struct {
			OneH float64 `json:"1h"`
		} `json:"rain,omitempty"`
	} `json:"hourly"`
	Daily []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Summary   string  `json:"summary"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Rain   float64 `json:"rain"`
		Uvi    float64 `json:"uvi"`
	} `json:"daily"`
}

func GetWeather() Useful {
	w := Weather{}
	u := Useful{}
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&appid=%s&units=imperial&exclude=minutely", LAT, LONG, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return u
	}
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &w)

	u.Current = fmt.Sprintf("%.3f", w.Current.Temp)
	u.High = fmt.Sprintf("%.3f", w.Daily[0].Temp.Max)
	u.Low = fmt.Sprintf("%.3f", w.Daily[0].Temp.Min)
	u.Percipitation = fmt.Sprintf("%.3f", w.Daily[0].Rain)
	u.Weather = strings.Replace(w.Current.Weather[0].Description, " ", "-", -1)
	fmt.Println(u)

	return u
}

func GetWeatherWithLatAndLong(lat float64, long float64) Useful {
	w := Weather{}
	u := Useful{}
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%.4f&lon=%.4f&appid=%s&units=imperial&exclude=minutely", lat, long, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return u
	}
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &w)

	u.Current = fmt.Sprintf("%.3f", w.Current.Temp)
	u.High = fmt.Sprintf("%.3f", w.Daily[0].Temp.Max)
	u.Low = fmt.Sprintf("%.3f", w.Daily[0].Temp.Min)
	u.Percipitation = fmt.Sprintf("%.3f", w.Daily[0].Rain)
	u.Weather = strings.Replace(w.Current.Weather[0].Description, " ", "-", -1)
	fmt.Println(u)

	return u
}

type LatAndLong struct {
	Zip     string
	Name    string
	Lat     float64
	Long    float64
	Country string
}

func GetLatAndLong(zip string) LatAndLong {
	l := LatAndLong{}
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s&limit=5&appid=%s", zip, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		return l
	}
	body, _ := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	err = json.Unmarshal(body, &l)
	if err != nil {
		fmt.Println(err)
		return l
	}
	return l
}
