package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ProtectedHandler(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get username"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "username": username})
}
