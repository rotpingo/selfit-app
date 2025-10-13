package routes

import (
	"selfit/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	handlers.RegisterAuthRoutes(server)
	handlers.RegisterNoteRoutes(server)
	handlers.RegisterTaskRoutes(server)
	handlers.RegisterTrackerRoutes(server)
	handlers.RegisterWeatherRoutes(server)
	handlers.RegisterUserRoutes(server)
}
