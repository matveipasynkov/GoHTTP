package weather_test

import (
	"GoHTTP/app/geo"
	"GoHTTP/app/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	//Arrange
	expected := "London"
	geodata := geo.GeoData{
		City: expected,
	}
	format := 3
	//Act
	result, _ := weather.GetWeather(geodata, format)
	//Assert
	if !strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, result)
	}
}

var testCases = []struct{
	name string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t* testing.T) {
			//Arrange
			expected := weather.ErrorIncorrectFormat
			city := "London"
			geodata := geo.GeoData{
				City: city,
			}
			//Act
			_, err := weather.GetWeather(geodata, tc.format)
			//Assert
			if err != expected {
				t.Errorf("Ожидалось %v, получение %v", expected, err)
			}
		})
	}
}