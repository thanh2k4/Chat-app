package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/thanh2k4/Chat-app/configs"
)

func GenerateToken(userId string, cfg configs.Config) (string, string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(cfg.JWT.RefreshTokenExpiry).Unix(),
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.JWT.SecretRefreshKey))
	if err != nil {
		return "", "", err
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(cfg.JWT.AccessTokenExpiry).Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(cfg.JWT.SecretAccessKey))
	if err != nil {
		return "", "", err
	}
	return refreshTokenString, accessTokenString, nil

}

func ValidateRefreshToken(tokenString string) (*jwt.Token, error) {

}
