package routes

import (
	"selfit/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	handlers.RegisterUserRoutes(server)
	handlers.RegisterNoteRoutes(server)
	handlers.RegisterTaskRoutes(server)
}
