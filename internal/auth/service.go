package auth

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/redis/go-redis/v9"
	"github.com/thanh2k4/Chat-app/internal/auth/infras/postgres"
	"github.com/thanh2k4/Chat-app/pkg/config"
	"github.com/thanh2k4/Chat-app/pkg/security"
	auth "github.com/thanh2k4/Chat-app/proto/gen"
	"time"
)

type AuthServer struct {
	auth.UnimplementedAuthServiceServer
	Postgres    *postgres.Queries
	RedisClient *redis.Client
	Cfg         config.Config
}

func (s *AuthServer) Refresh(ctx context.Context, req *auth.RefreshRequest) (*auth.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	claims, err := security.ValidateRefreshToken(req.RefreshToken, s.Cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %v", err)
	}

	userID := (*claims)["userId"].(string)
	storedToken, err := s.RedisClient.Get(ctx, userID).Result()
	if err != nil {
		return nil, fmt.Errorf("refresh token not found in Redis: %v", err)
	}

	if storedToken != req.RefreshToken {
		return nil, fmt.Errorf("refresh token does not match")
	}

	refreshToken, accessToken, err := security.GenerateToken(uuid.MustParse(userID), s.Cfg)
	if err != nil {
		return nil, fmt.Errorf("could not generate tokens: %v", err)
	}

	err = s.RedisClient.Set(ctx, userID, refreshToken, s.Cfg.JWT.RefreshTokenExpiry).Err()
	if err != nil {
		return nil, fmt.Errorf("could not store refresh token in Redis: %v", err)
	}

	return &auth.AuthResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (s *AuthServer) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hashedPassword, err := security.Hash(req.Password)
	if err != nil {
		return nil, fmt.Errorf("could not hash password: %v", err)
	}

	userID := uuid.New()

	_, err = s.Postgres.GetUserByUsername(ctx, req.Username)
	if err == nil {
		return nil, fmt.Errorf("user with username %s already exists", req.Username)
	}

	_, err = s.Postgres.CreateUser(ctx, postgres.CreateUserParams{
		ID:       pgtype.UUID{Bytes: userID, Valid: true},
		Username: req.Username,
		Password: string(hashedPassword),
	})

	if err != nil {
		return nil, fmt.Errorf("could not create user in database: %v", err)
	}

	refreshToken, accessToken, err := security.GenerateToken(userID, s.Cfg)
	if err != nil {
		return nil, fmt.Errorf("could not generate tokens: %v", err)
	}

	err = s.RedisClient.Set(ctx, userID.String(), refreshToken, s.Cfg.JWT.RefreshTokenExpiry).Err()
	if err != nil {
		return nil, fmt.Errorf("could not store refresh token in Redis: %v", err)
	}

	return &auth.AuthResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (s *AuthServer) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	user, err := s.Postgres.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}

	err = security.VerifyPassword(user.Password, req.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid password: %v", err)
	}

	refreshToken, accessToken, err := security.GenerateToken(uuid.UUID(user.ID.Bytes), s.Cfg)
	if err != nil {
		return nil, fmt.Errorf("could not generate tokens: %v", err)
	}

	err = s.RedisClient.Set(ctx, user.ID.String(), refreshToken, s.Cfg.JWT.RefreshTokenExpiry).Err()
	if err != nil {
		return nil, fmt.Errorf("could not store refresh token in Redis: %v", err)
	}

	return &auth.AuthResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
