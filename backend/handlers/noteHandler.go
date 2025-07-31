package handlers

import (
	"fmt"
	"net/http"
	"selfit/models"
	"selfit/services"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(server *gin.Engine) {

	server.GET("/api/notes", getNotes)
	server.POST("/api/notes", createNote)

}

func getNotes(context *gin.Context) {
	notes, err := services.GetAllNotes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch notes. Try again later."})
		return
	}
	context.JSON(http.StatusOK, notes)
}

// TODO: Finish this method
func createNote(context *gin.Context) {
	var note models.Note
	err := context.ShouldBindJSON(&note)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	fmt.Println(note)

	err = services.SaveNote(note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Note. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Note Created!", "note": note})
}
