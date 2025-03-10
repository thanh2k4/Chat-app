package auth

import (
	"context"
	"github.com/thanh2k4/Chat-app/cmd/auth/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/pkg/security"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RefreshHandler(cfg config.Config, redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RefreshRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		claims, err := security.ValidateRefreshToken(req.RefreshToken, cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}
		ctx := context.Background()
		storedToken, err := redisClient.Get(ctx, (*claims)["userId"].(string)).Result()
		if err != nil || storedToken != req.RefreshToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}

		refreshToken, accessToken, err := security.GenerateToken((*claims)["userId"].(string), cfg)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
			return
		}

		redisClient.Set(ctx, (*claims)["userId"].(string), refreshToken, cfg.JWT.RefreshTokenExpiry)
		c.JSON(http.StatusOK, gin.H{
			"refresh_token": refreshToken,
			"access_token":  accessToken,
		})

	}
}

func LoginHandler(cfg config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//user, err := AuthenticatedUser(req.Username, req.Password)
		//if err != nil {
		//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not authenticate user"})
		//	return
		//}
	}
}
