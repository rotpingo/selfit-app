package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/services"
	"selfit/utils"

	"github.com/gin-gonic/gin"
)

func RegisterWeatherRoutes(server *gin.Engine) {
	server.POST("/api/weather", addCity)
}

// func fetchData(context *gin.Context) {
//
// 	var req dto.WeatherRequestDTO
// 	err := context.ShouldBindJSON(&req)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	weatherDto, err := services.FetchWeather(req.City)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	context.JSON(http.StatusOK, weatherDto)
// }

// func searchCity(context *gin.Context) {
// 	utils.CheckToken(context)
// 	token := context.Request.Header.Get("Authorization")
// 	userId, err := utils.VerifyToken(token)
// 	if err != nil {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
// 		return
// 	}
// 	city := context.Query("city")
// 	if city == "" {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": "city query parameter is required"})
// 		return
// 	}
// }

func addCity(context *gin.Context) {

	utils.CheckToken(context)
	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var req dto.WeatherRequestDTO
	err = context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weatherDto, err := services.FetchCity(req.City)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	weather := weatherDto.ToWeatherModel(userId)
	services.CreateWeather(weather)

	context.JSON(http.StatusOK, gin.H{"message": "City successfuly added"})
}
