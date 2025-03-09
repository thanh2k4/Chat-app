package security

import (
	"fmt"
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

func ValidateRefreshToken(tokenString string, cfg configs.Config) (*jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.SecretRefreshKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}
	return nil, nil

}
