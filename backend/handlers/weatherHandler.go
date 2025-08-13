package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/services"
	"selfit/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterWeatherRoutes(server *gin.Engine) {
	server.GET("/api/weather", getCities)
	server.POST("/api/weather", addCity)
	server.GET("/api/weather/:id", fetchCityData)
	server.DELETE("/api/weather/:id", deleteCity)
}

func getCities(context *gin.Context) {

	utils.CheckToken(context)
	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	cities, err := services.GetAllCities(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch notes. Try again later."})
		return
	}

	context.JSON(http.StatusOK, cities)
}

func fetchCityData(context *gin.Context) {

	utils.CheckToken(context)
	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	cityId := context.Param("id")
	intId, err := strconv.Atoi(cityId)
	id := int64(intId)

	cityData, err := services.FetchWeather(userId, id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"API error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cityData)
}

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

	weatherDto, err := services.FetchCity(req.Name)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	weather := weatherDto.ToWeatherModel(userId)
	services.CreateWeather(weather)

	context.JSON(http.StatusOK, gin.H{"message": "City successfuly added"})
}

func deleteCity(context *gin.Context) {

	utils.CheckToken(context)
	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	cityId := context.Param("id")
	intId, err := strconv.Atoi(cityId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid city ID"})
		return
	}

	err = services.DeleteCityById(id, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "City deleted successfully!"})
}
