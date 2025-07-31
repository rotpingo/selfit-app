package controllers

import (
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/api/notes", func(context *gin.Context) {
		notes := models.GetAllNotes()
		context.JSON(http.StatusOK, notes)
	})

}
