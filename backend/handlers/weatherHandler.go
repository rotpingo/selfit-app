package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/services"

	"github.com/gin-gonic/gin"
)

func RegisterWeatherRoutes(server *gin.Engine) {
	server.GET("/api/weather", fetchData)
}

func fetchData(context *gin.Context) {

	var req dto.WeatherRequestDTO
	err := context.ShouldBindJSON(&req)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weatherDto, err := services.FetchWeather(req.City)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, weatherDto)
}
