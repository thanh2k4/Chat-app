package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/configs"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func RefreshHandler(cfg configs.Config, redisClient *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req RefreshRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid Refresh Token"})
			return
		}
	}
}
