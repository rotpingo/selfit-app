package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/services"
	"selfit/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(server *gin.Engine) {
	server.GET("api/user", getUser)
}

func getUser(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch user data. Try again later."})
		return
	}

	// workaround for parsing null to string
	userDto := dto.UserResponseDTO{
		ID:    user.ID,
		Name:  "",
		Email: user.Email,
	}

	// verify if the value is not null
	if user.Name.Valid {
		userDto.Name = user.Name.String
	}
	context.JSON(http.StatusOK, userDto)
}
