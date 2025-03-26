package handler

import (
	"github.com/gin-gonic/gin"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/proto/gen"
)

func CreateUserHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.UserClient.CreateUser(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func GetUserHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.GetUserRequest
		req.Id = c.GetString("userId")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.UserClient.GetUser(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}

func UpdateUserHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.UpdateUserRequest
		req.Id = c.GetString("userId")
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.UserClient.UpdateUser(c, &req)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, resp)
	}
}
