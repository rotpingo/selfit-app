package handlers

import (
	"net/http"
	"selfit/models"
	"selfit/services"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("/auth/login", login)
	server.POST("/auth/register", register)
}

func register(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = services.CreateUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create User. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User Created!", "user": user})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = services.ValidateUser(&user)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
