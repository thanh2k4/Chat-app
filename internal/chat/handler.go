package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/thanh2k4/Chat-app/internal/chat/infras/postgres"
	"net/http"
)

func SendMessageHandler(manager *Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req postgres.ChatMessage
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := PublishMessage(manager.redisClient, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
