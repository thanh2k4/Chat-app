package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	auth "github.com/thanh2k4/Chat-app/proto/gen"
	"google.golang.org/grpc"
	"net/http"
)

type AuthHandler struct {
	authClient auth.AuthServiceClient
}

func NewAuthHandler(authConn *grpc.ClientConn) *AuthHandler {
	return &AuthHandler{
		authClient: auth.NewAuthServiceClient(authConn),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req auth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.authClient.RegisterHandler(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req auth.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.authClient.RefreshHandler(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.authClient.LoginHandler(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
