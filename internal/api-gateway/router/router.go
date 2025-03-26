package router

import (
	"github.com/gin-gonic/gin"
	apigateway "github.com/thanh2k4/Chat-app/internal/api-gateway"
	"github.com/thanh2k4/Chat-app/internal/api-gateway/handler"
	"github.com/thanh2k4/Chat-app/internal/api-gateway/middleware"
)

func SetupAPIGatewayRoutes(client *apigateway.Client) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.AuthMiddleware(client))
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.RequestLoggerMiddleware())

	// Auth Routes
	r.POST("/auth/register", handler.RegisterHandler(client))
	r.POST("/auth/login", handler.LoginHandler(client))
	r.POST("/auth/refresh", handler.RefreshHandler(client))
	r.POST("/auth/logout", handler.LogoutHandler(client))

	// Chat Routes
	r.GET("/chats/:id", handler.GetChatHandler(client))
	r.POST("/chats", handler.CreateChatHandler(client))
	r.GET("/chats/user", handler.GetChatByUserHandler(client))
	r.POST("/chats/:id/messages", handler.SendMessageHandler(client))
	r.GET("/chats/:id/messages", handler.GetMessageByChat(client))

	// User Routes
	r.POST("/users", handler.CreateUserHandler(client))
	r.GET("/users", handler.GetUserHandler(client))
	r.PUT("/users", handler.UpdateUserHandler(client))

	return r
}
