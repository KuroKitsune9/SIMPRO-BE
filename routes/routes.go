package routes

import (
	"simpro/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.GET("/users", handler.GetUsers)
	}
}
