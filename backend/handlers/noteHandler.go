package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/models"
	"selfit/services"
	"selfit/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterNoteRoutes(server *gin.Engine) {
	server.GET("/api/notes", getNotes)
	server.POST("/api/notes", createNote)
	server.DELETE("/api/notes/:id", deleteNote)
	server.PUT("/api/notes/:id", updateNote)
}

func getNotes(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	notes, err := services.GetAllNotes(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch notes. Try again later."})
		return
	}
	context.JSON(http.StatusOK, notes)
}

func createNote(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var noteDto dto.CreateNoteDTO
	err = context.ShouldBindJSON(&noteDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	note := noteDto.ToNoteModel(userId)
	err = services.CreateNote(note)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Note. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Note Created!", "note": note})
}

func updateNote(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var note models.Note
	err = context.ShouldBindJSON(&note)

	err = services.UpdateNote(&note, userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note modified!", "note": note})
}

func deleteNote(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	noteId := context.Param("id")
	intId, err := strconv.Atoi(noteId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	err = services.DeleteNoteById(id, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully!"})
}
