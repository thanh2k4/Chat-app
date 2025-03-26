package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/pkg/security"
	"github.com/thanh2k4/Chat-app/proto/gen"
	"net/http"
	"time"
)

var publicRoutes = map[string]bool{
	"/auth/register": true,
	"/auth/login":    true,
}

func AuthMiddleware(client *apigateway.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := publicRoutes[c.Request.URL.Path]; ok {
			c.Next()
			return
		}
		accessToken, err := c.Cookie("access_token")
		cfg, err := config.LoadConfig("cmd/auth/config/config.yaml")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
		}
		if err != nil {
			claims, err := security.ValidateAccessToken(accessToken, *cfg)
			if err == nil {
				userID := (*claims)["userId"].(string)
				c.Set("userID", userID)
				c.Next()
				return
			}
		}
		var req *gen.RefreshRequest
		if err := c.ShouldBindJSON(req); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		claims, err := security.ValidateRefreshToken(req.RefreshToken, *cfg)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		resp, err := client.AuthClient.Refresh(ctx, req)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.SetCookie("access_token", resp.AccessToken, 900, "/", "localhost", true, true)
		c.Set("UserId", (*claims)["userId"].(string))
		c.Set("refresh_token", req.RefreshToken)
		c.Next()
	}
}
