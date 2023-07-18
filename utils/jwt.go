package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
)

func GetToken(username, jwtSecret string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	return token.SignedString([]byte(jwtSecret))
}

func GetJWTSecret() (string, error) {
	jwtSecret, exists := os.LookupEnv("JWT_SECRET")
	if !exists {
		return "", errors.New("environment variable JWT_SECRET not set")
	}

	return jwtSecret, nil
}
