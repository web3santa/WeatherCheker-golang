package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
	// Location struct {
	// 	Lat string `json:"latitude"`
	// 	Lng string `json:"longitude"`
	// } `json:"location"`
	Current struct {
		Time       string  `json:"time"`
		Temperture float64 `json:"temperature_2m"`
	} `json:"current"`
}

// https://open-meteo.com/ from free api for non commercial use only
func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m"
	// 서울 은평구 경도 위도
	// latitude := "126.9312417"
	// longitude := "37.59996944"

	// url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=126.9312417&longitude=37.59996944&current=temperature_2m,wind_speed_10m&hourly=temperature_2m,relative_humidity_2m,wind_speed_10m")
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

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}
	fmt.Println(weather)
}
