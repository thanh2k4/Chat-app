package auth

import "github.com/google/uuid"

type UserRegister struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" binding:"required,min=5,max=32"`
	Password string    `json:"password" binding:"required,min=8,max=32"`
}

type UserLogin struct {
	Username string `json:"username"  binding:"required,min=5,max=32"`
	Password string `json:"password" binding:"required,min=8,max=32"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
