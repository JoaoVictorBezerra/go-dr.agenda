package routes

import (
	"dr.agenda/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, handler *handler.UserHandler) {
	user := r.Group("/api/user")
	{
		user.GET("", handler.GetUsers)
	}
}
