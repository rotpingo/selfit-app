package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"selfit/dto"
)

func FetchWeather(city string) (dto.WeatherDTO, error) {

	apiKey := getApiKey()
	apiUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city,
		apiKey,
	)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return dto.WeatherDTO{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return dto.WeatherDTO{}, fmt.Errorf("API error: %s", string(bodyBytes))
	}
	var weatherDto dto.WeatherDTO
	if err := json.NewDecoder(resp.Body).Decode(&weatherDto); err != nil {
		return dto.WeatherDTO{}, fmt.Errorf("failed to parse response: %w", err)
	}

	return weatherDto, nil
}

func getApiKey() string {
	return os.Getenv("API_WEATHER_KEY")
}
