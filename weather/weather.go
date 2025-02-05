package weather

import (
	"GoHTTP/app/geo"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrorIncorrectFormat = errors.New("INCORRECT_FORMAT")

func GetWeather(geo geo.GeoData, format int) (string, error) {
	if format > 4 || format < 1 {
		fmt.Println(ErrorIncorrectFormat.Error())
		return "", ErrorIncorrectFormat
	}
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	return string(body), nil
}