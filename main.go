package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Структуры для парсинга данных из API
type WeatherResponse struct {
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
	Timezone string  `json:"timezone"`
	Current  Current `json:"current"`
	Daily    []Daily `json:"daily"`
}

type Current struct {
	Temp      float64   `json:"temp"`
	FeelsLike float64   `json:"feels_like"`
	Weather   []Weather `json:"weather"`
}

type Daily struct {
	Temp    Temp      `json:"temp"`
	Weather []Weather `json:"weather"`
}

type Temp struct {
	Day float64 `json:"day"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

type Weather struct {
	Description string `json:"description"`
}

func main() {
	apiKey := "76f2a967651d25b1e25a69fa6058be83"
	lat := "69.7671"
	lon := "21.0256"

	// Формирование URL запроса
	url := fmt.Sprintf("https://api.openweathermap.org/data/3.0/onecall?lat=%s&lon=%s&appid=%s&units=metric", lat, lon, apiKey)

	// Отправка GET запроса
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Декодируем ответ в структуру WeatherResponse
	var weatherResponse WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		fmt.Println("Error while decoding response:", err)
		os.Exit(1)
	}

	// Пример вывода текущей погоды
	fmt.Printf("Current temperature: %.2f°C\n", weatherResponse.Current.Temp)
	fmt.Printf("Feels like: %.2f°C\n", weatherResponse.Current.FeelsLike)
	fmt.Printf("Description: %s\n", weatherResponse.Current.Weather[0].Description)

	// Переменные для хранения минимальных и максимальных температур за несколько дней
	dailyMinTemp := weatherResponse.Daily[0].Temp.Min
	dailyMaxTemp := weatherResponse.Daily[0].Temp.Max

	// Пример вывода прогноза на каждый день
	fmt.Println("\nDaily Forecast:")
	for _, daily := range weatherResponse.Daily {
		// Обновление минимальной и максимальной температуры
		if daily.Temp.Min < dailyMinTemp {
			dailyMinTemp = daily.Temp.Min
		}
		if daily.Temp.Max > dailyMaxTemp {
			dailyMaxTemp = daily.Temp.Max
		}

		// Вывод прогноза на день без округления
		fmt.Printf("Day Temp: %.2f°C, Min: %.2f°C, Max: %.2f°C\n", daily.Temp.Day, daily.Temp.Min, daily.Temp.Max)
		fmt.Printf("Description: %s\n", daily.Weather[0].Description)
		fmt.Println("-----------")
	}

	// Вывод минимальной и максимальной температуры за несколько дней
	fmt.Printf("\nOverall Min Temperature: %.2f°C\n", dailyMinTemp)
	fmt.Printf("Overall Max Temperature: %.2f°C\n", dailyMaxTemp)
}
