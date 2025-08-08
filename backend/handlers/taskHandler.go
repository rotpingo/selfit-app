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

func RegisterTaskRoutes(server *gin.Engine) {
	server.GET("/api/tasks", getTasks)
	server.POST("/api/tasks", createTask)
	server.DELETE("/api/tasks/:id", deleteTask)
	server.PUT("/api/tasks/:id", updateTask)
	server.PATCH("/api/tasks/:id/abort", abortTask)
	server.PATCH("/api/tasks/:id/complete", completeTask)
}

func getTasks(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	tasks, err := services.GetAllProgressTasks(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch tasks. Try again later."})
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func createTask(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var task models.Task
	err = context.ShouldBindJSON(&task)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = services.CreateTask(&task, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Task. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Task Created!", "task": task})
}

func updateTask(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var task models.Task
	err = context.ShouldBindJSON(&task)

	err = services.UpdateTask(&task, userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Note modified!", "task": task})
}

func deleteTask(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	taskId := context.Param("id")
	intId, err := strconv.Atoi(taskId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = services.DeleteTaskById(id, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
}

func abortTask(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	taskId := context.Param("id")
	intId, err := strconv.Atoi(taskId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var taskDto dto.EndTaskDTO
	err = context.ShouldBindJSON(&taskDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body for task notes"})
		return
	}
	err = services.AbortTaskById(id, taskDto, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not abort Task. Try again Later."})
		return
	}
}

func completeTask(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	taskId := context.Param("id")
	intId, err := strconv.Atoi(taskId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}
	var taskDto dto.EndTaskDTO
	err = context.ShouldBindJSON(&taskDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body for task notes"})
		return
	}
	err = services.CompleteTaskById(id, taskDto, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not complete Task. Try again Later."})
		return
	}
}
