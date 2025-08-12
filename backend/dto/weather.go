package dto

import ()

type WeatherRequestDTO struct {
	City string `json:"city" binding:"required"`
}

type WeatherDTO struct {
	Name string `json:"name"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Sys struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Weather []struct {
		Id          int64  `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}
}

// func (dto WeatherDTO) ToWeatherModel(userId int64) *models.Weather {
// 	return &models.Weather{
// 		City:        dto.City,
// 		Country:     dto.Country,
// 		Temperature: dto.Temperature,
// 		Description: dto.Description,
// 		Humidity:    dto.Humidity,
// 		Wind:        dto.Wind,
// 		Cloudiness:  dto.Cloudiness,
// 		WeatherCode: dto.WeatherCode,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 		UserID:      userId,
// 	}
// }
//
// func WeatherToResponseDTO(weather *models.Weather) WeatherDTO {
// 	return WeatherDTO{
// 		ID:          weather.ID,
// 		City:        weather.City,
// 		Country:     weather.Country,
// 		Temperature: weather.Temperature,
// 		Description: weather.Description,
// 		Humidity:    weather.Humidity,
// 		Wind:        weather.Wind,
// 		Cloudiness:  weather.Cloudiness,
// 		WeatherCode: weather.WeatherCode,
// 	}
// }
