package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weatherGo/geo"
)

var ErrWrongFormat = errors.New("wrong format")
var ErrUrl = errors.New("error_url")
var ErrHttp = errors.New("error_http")
var ErrReadBody = errors.New("error_readBody")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format < 1 || format > 4 {
		return "", ErrWrongFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrUrl
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrHttp
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", ErrReadBody
	}
	return string(body), nil
}
