package geo_test

import (
	"testing"
	"weatherGo/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "London"
	expectedCity := geo.GeoData{
		City: "London",
	}
	// Act
	got, err := geo.GetMyLocation(city)
	// Assert
	if err != nil {
		t.Error(err)
	}
	if got.City != expectedCity.City {
		t.Errorf("Ожидалось %v, получено %v", expectedCity.City, got.City)
	}
}

func TestGetMyLocationInvalidCity(t *testing.T) {
	// Arrange
	city := "Lond"
	// Act
	_, err := geo.GetMyLocation(city)
	// Assert
	if err != geo.ErrInvalidCity {
		t.Errorf("Ожидалось %v, получено %v", geo.ErrInvalidCity, err)
	}
}
