package chat

import "github.com/gin-gonic/gin"

func SetupChatRoutes(manager *Manager) *gin.Engine {
	r := gin.Default()
	r.GET("/ws", manager.HandleConnection)
	r.POST("/messages", SendMessageHandler(manager))
	return r
}
