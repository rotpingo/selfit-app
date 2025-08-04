package handlers

import (
	"fmt"
	"net/http"
	"selfit/models"
	"selfit/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(server *gin.Engine) {

	server.GET("/api/notes", getNotes)
	server.POST("/api/notes", createNote)
	server.DELETE("/api/notes/:id", deleteNote)
	server.PUT("/api/notes/:id", editNote)

}

func getNotes(context *gin.Context) {
	notes, err := services.GetAllNotes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch notes. Try again later."})
		return
	}
	context.JSON(http.StatusOK, notes)
}

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

func editNote(context *gin.Context) {

	var note models.Note
	err := context.ShouldBindJSON(&note)

	err = services.EditNote(note)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note modified!", "note": note})
}

func deleteNote(context *gin.Context) {
	noteId := context.Param("id")

	id, err := strconv.Atoi(noteId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	err = services.DeleteNoteById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully!"})
}
