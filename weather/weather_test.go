package weather_test

import (
	"strings"
	"testing"
	"weatherGo/geo"
	"weatherGo/weather"
)

func TestGetWeather(t *testing.T) {
	// Arrange
	format := 4
	expectedCity := "Madrid"
	geoData := geo.GeoData{
		City: expectedCity,
	}
	// Act
	result, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	// Assert
	if !strings.Contains(result, expectedCity) {
		t.Errorf("Ожидалось %v, получено %v", expectedCity, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 150},
	{name: "Zero format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expectedCity := "Madrid"
			geoData := geo.GeoData{
				City: expectedCity,
			}
			_, err := weather.GetWeather(geoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, получено %v", weather.ErrWrongFormat, err)
			}
		})
	}
}
