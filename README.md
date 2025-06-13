# Pico Weather Display Backend
This is the backend to the [pico-weather-display](https://github.com/mmnessim/pico-weather-display.git) project. It accepts requests from the Pico and returns weather data formatted for parsing by the Pico to display.

## Getting Started
First, make sure you have [Go](https://go.dev/doc/install) installed. Then clone the repository:
```bash
git clone https://github.com/mmnessim/pico-weather-display-backend.git
cd pico-weather-display-backend
```
Create a .env file for your OpenWeatherMap API Key in the root of the directory
```bash
touch .env
echo "API_KEY=" > .env
```
Visit [OpenWeatherMap](https://openweathermap.org/api) to create and account and get a free API key, then paste it into your .env file. Then simple run:
```bash
go run .
```

