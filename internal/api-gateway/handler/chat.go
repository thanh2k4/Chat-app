package handler

import (
	"github.com/gin-gonic/gin"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/proto/gen"
)

func SendMessageHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.ChatMessage
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ChatClient.SendMessage(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func CreateChatHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.ChatCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ChatClient.CreateChat(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func GetChatHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.ChatGetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ChatClient.GetChat(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func GetChatByUserHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.ChatGetByUserRequest
		req.UserId = c.GetString("userId")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ChatClient.GetChatsByUser(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func GetMessageByChat(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.ChatGetRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.ChatClient.GetMessagesByChat(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}
