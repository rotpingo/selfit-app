package handlers

import (
	"net/http"
	"selfit/models"
	"selfit/services"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(server *gin.Engine) {
	server.GET("/api/tasks", getTasks)
	server.POST("/api/tasks", createTask)
	// server.DELETE("/api/tasks/:id", deleteTask)
	// server.PUT("/api/tasks/:id", editTask)
}

func getTasks(context *gin.Context) {

	tasks, err := services.GetAllTasks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	var task models.Task
	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = services.SaveTask(task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Task. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task Created!", "task": task})
}
