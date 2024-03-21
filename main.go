package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Weather struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`

	Current struct {
		Time       string  `json:"time"`
		Temperture float64 `json:"temperature_2m"`
	} `json:"current"`
}

// https://open-meteo.com/ from free api for non commercial use only
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/weather", Weathers)

	http.ListenAndServe(":8080", r)
}

func Weathers(w http.ResponseWriter, r *http.Request) {
	url := "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(body))
	var weather Weather // 포인터로 선언

	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	coordinates := fmt.Sprintf("위도: %f, 경도: %f", weather.Lat, weather.Lng)

	time, temp := weather.Current.Time, weather.Current.Temperture

	message := fmt.Sprintf("location: %s, time: %s, temp: %f", coordinates, time, temp)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, message)
}
