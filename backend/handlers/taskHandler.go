package handlers

import (
	"fmt"
	"net/http"
	"selfit/models"
	"selfit/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterTaskRoutes(server *gin.Engine) {
	server.GET("/api/tasks", getTasks)
	server.POST("/api/tasks", createTask)
	server.DELETE("/api/tasks/:id", deleteTask)
	server.PUT("/api/tasks/:id", updateTask)
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
	fmt.Println(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = services.CreateTask(&task)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Task. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task Created!", "task": task})
}

func updateTask(context *gin.Context) {

	var task models.Task
	err := context.ShouldBindJSON(&task)

	err = services.UpdateTask(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note modified!", "task": task})
}

func deleteTask(context *gin.Context) {
	taskId := context.Param("id")
	id, err := strconv.Atoi(taskId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = services.DeleteTaskById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}
