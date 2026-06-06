package handler

import (
	"net/http"
	"simpro/internal/service"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := service.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}
