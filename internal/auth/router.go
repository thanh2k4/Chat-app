package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/configs"
)

type AuthRouter struct {
	Config      configs.Config
	redisClient *redis.Client
}

func (r *AuthRouter) RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	routes.POST("/refresh", RefreshHandler(r.Config, r.redisClient))

}
