package dto

import (
	"selfit/models"
	"time"
)

type WeatherRequestDTO struct {
	City string `json:"city" binding:"required"`
}

type WeatherAPIResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
	}
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}
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

type CreateWeatherDTO struct {
	Name    string  `json:"name"`
	Country string  `json:"country"`
	Lon     float64 `json:"lon"`
	Lat     float64 `json:"lat"`
}

func (dto CreateWeatherDTO) ToWeatherModel(userId int64) *models.Weather {
	return &models.Weather{
		Name:      dto.Name,
		Country:   dto.Country,
		Lon:       dto.Lon,
		Lat:       dto.Lat,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userId,
	}
}

func NewCreateWeatherDTOFromAPI(api WeatherAPIResponse) CreateWeatherDTO {
	return CreateWeatherDTO{
		Name:    api.Name,
		Country: api.Sys.Country,
		Lon:     api.Coord.Lon,
		Lat:     api.Coord.Lat,
	}
}
