package geo_test

import (
	"GoHTTP/app/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	//Arrange - подготовка, expected result, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}
	//Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	//Assert - проверка результата с expected
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "LOdasdsdfdsfs"

	_, err := geo.GetMyLocation(city)
	
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}
}