package routes

import (
	"simpro/internal/handler"
	"simpro/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	authProtected := api.Group("/auth")
	authProtected.Use(middleware.JWTMiddleware())
	{
		authProtected.GET("/profile", handler.Me)
	}
}
