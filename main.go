package main

import (
	"flag"
	"fmt"
	"weatherGo/geo"
	"weatherGo/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "Формат вывода погоды")

	flag.Parse()

	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	weatherData, _ := weather.GetWeather(*geoData, *format)
	fmt.Println(weatherData)
}
