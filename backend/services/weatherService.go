package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"selfit/database"
	"selfit/dto"
	"selfit/models"
)

func GetAllCities(userId int64) ([]dto.WeatherResponseDTO, error) {
	query := "SELECT id, name, country FROM weather WHERE user_id = $1"
	rows, err := database.DB.Query(query, userId)
	if err != nil {
		fmt.Println("error fetching:", err)
		return nil, err
	}
	defer rows.Close()

	var cities []dto.WeatherResponseDTO
	for rows.Next() {
		var cityData dto.WeatherResponseDTO
		err := rows.Scan(&cityData.ID, &cityData.Name, &cityData.Country)
		if err != nil {
			return nil, err
		}
		cities = append(cities, cityData)
	}
	return cities, nil
}
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

func FetchCity(city string) (dto.CreateWeatherDTO, error) {

	apiKey := getApiKey()
	apiUrl := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric",
		city,
		apiKey,
	)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return dto.CreateWeatherDTO{}, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return dto.CreateWeatherDTO{}, fmt.Errorf("API error: %s", string(bodyBytes))
	}
	var weatherAPI dto.WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherAPI); err != nil {
		return dto.CreateWeatherDTO{}, fmt.Errorf("failed to parse response: %w", err)
	}

	weatherDto := dto.NewCreateWeatherDTOFromAPI(weatherAPI)

	return weatherDto, nil
}

func CreateWeather(weather *models.Weather) error {

	query := `
	INSERT INTO weather(name, country, lon, lat, created_at, updated_at, user_id) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	RETURNING id
	`

	err := database.DB.QueryRow(
		query,
		weather.Name,
		weather.Country,
		weather.Lon,
		weather.Lat,
		weather.CreatedAt,
		weather.UpdatedAt,
		weather.UserID,
	).Scan(&weather.ID)

	if err != nil {
		fmt.Println("insert error:", err)
		return err
	}

	return nil
}

func getApiKey() string {
	return os.Getenv("API_WEATHER_KEY")
}
