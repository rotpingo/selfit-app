package handlers

import (
	"net/http"
	"selfit/dto"
	"selfit/services"
	"selfit/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterTrackerRoutes(server *gin.Engine) {
	server.GET("/api/trackers", getTrackers)
	server.POST("/api/trackers", createTracker)
	server.PUT("/api/trackers/:id", updateTracker)
	server.DELETE("/api/trackers/:id", deleteTracker)
}

func getTrackers(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	trackers, err := services.GetAllTrackers(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch trackers. Try again later."})
		return
	}

	var trackersDto []dto.TrackerResponseDTO
	for _, tracker := range trackers {
		trackersDto = append(trackersDto, dto.TrackerToResponseDTO(&tracker))
	}

	context.JSON(http.StatusOK, trackersDto)
}

func createTracker(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var trackerDto dto.CreateTrackerDTO
	err = context.ShouldBindJSON(&trackerDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	tracker := trackerDto.ToTrackerModel(userId)

	err = services.CreateTracker(tracker)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create Tracker. Try again Later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Tracker Created!"})
}

func updateTracker(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	var trackerDto dto.UpdateTrackerDTO
	err = context.ShouldBindJSON(&trackerDto)

	tracker := trackerDto.ToTrackerModel(userId)
	err = services.UpdateTracker(tracker)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Tracker modified!"})
}

func deleteTracker(context *gin.Context) {
	utils.CheckToken(context)

	token := context.Request.Header.Get("Authorization")
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	trackerId := context.Param("id")
	intId, err := strconv.Atoi(trackerId)
	id := int64(intId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid note ID"})
		return
	}

	err = services.DeleteTrackerById(id, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Tracker deleted successfully!"})
}
