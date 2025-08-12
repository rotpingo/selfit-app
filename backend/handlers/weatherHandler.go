package handlers

import (
	"net/http"
	"selfit/services"

	"github.com/gin-gonic/gin"
)

func RegisterWeatherRoutes(server *gin.Engine) {
	server.GET("/api/weather", fetchData)
}

func fetchData(context *gin.Context) {

	city := context.Query("city")
	if city == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "city query parameter is required"})
		return
	}

	weatherDto, err := services.FetchWeather(city)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, weatherDto)
}
