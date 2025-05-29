package routes

import (
	"dr.agenda/handler"
	"github.com/gin-gonic/gin"
)

func RegisterAuthenticationRoutes(r *gin.Engine, handler *handler.AuthenticationHandler) {
	user := r.Group("/api/auth")
	{
		user.POST("/login", handler.Login)
	}
}
