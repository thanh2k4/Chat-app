package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"net/http"
)

func RegisterHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.RegisterRequest
		req.Id = uuid.New().String()
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.AuthClient.Register(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("access_token", resp.AccessToken, 900, "/", "localhost", true, true)
		c.JSON(http.StatusOK, resp)
	}
}

func LoginHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.AuthClient.Login(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("access_token", resp.AccessToken, 900, "/", "localhost", true, true)
		c.JSON(http.StatusOK, resp)
	}
}

func RefreshHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req gen.RefreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := client.Validator.Validate(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resp, err := client.AuthClient.Refresh(context.Background(), &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.SetCookie("access_token", resp.AccessToken, 900, "/", "localhost", true, true)
		c.JSON(http.StatusOK, resp)
	}
}

func LogoutHandler(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("access_token", "", -1, "/", "localhost", true, true)

		_, err := client.AuthClient.Logout(context.Background(), &gen.LogoutRequest{
			Id: c.GetString("userID"),
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "logged out"})
	}
}
