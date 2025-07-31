package main

import (
	"gin/config"
	"gin/database"
	"gin/models"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config.LoadEnv()
	database.Connect()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	server.GET("/api/notes", func(context *gin.Context) {
		notes := models.GetAllNotes()
		context.JSON(http.StatusOK, notes)
	})

	server.POST("/api/notes", createNote)

	// Start server on port 6969
	server.Run(":6969")
}

func createNote(context *gin.Context) {
	var note models.Note
	err := context.ShouldBindJSON(&note)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse Data"})
		return
	}

	note.ID = 1
	note.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "Note Created", "note": note})
}
