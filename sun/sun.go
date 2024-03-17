package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	q := "Tashkent"

	if len(os.Args) >= 2 {
		q = os.Args[1]
	}
	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=ec180872243c4f57a4f153631230105&q=" + q + "&days=1&aqi=no&alert=no")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Not available sorry... 😭")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}
		switch hour.Condition.Text {
		case "Sunny":
			hour.Condition.Text += " 🌞"
		case "Cloudy", "Partly Cloudy", "Overcast":
			hour.Condition.Text += " 🌥️"
		case "Clear":
			hour.Condition.Text += " 🌝"
		case "Light rain", "Light drizzle":
			hour.Condition.Text += " 🌦️"
		case "Rainy":
			hour.Condition.Text += " 🌧️"
		default:
			hour.Condition.Text += " 🌦️"
		}
		message := fmt.Sprintf(
			"%s -> %.0fC, %.0f%%, %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)

		if hour.TempC < 1.0 {
			color.Blue(message)
		} else if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		}  else {
			color.Red(message)
			
		}
	}
}