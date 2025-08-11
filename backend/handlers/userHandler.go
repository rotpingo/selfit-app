package handlers

import (
	"fmt"
	"net/http"
	"selfit/models"
	"selfit/services"
	"selfit/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.POST("api/auth/login", login)
	server.POST("api/auth/register", register)
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

	context.JSON(http.StatusCreated, gin.H{"message": "User Created!"})
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
		fmt.Println("you are here")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
