package util

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(username, tokenSymmetricKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["iss"] = "squrl"
	claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	tokenString, err := token.SignedString([]byte(tokenSymmetricKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(bearerToken, tokenSymmetricKey string) (string, error) {
	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return "", errors.New("Invalid bearer header, should start with 'Bearer '")
	}
	jwtToken := strings.Split(bearerToken, " ")[1]
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSymmetricKey), nil
	})
	if err != nil {
		return "", errors.New(fmt.Sprintln(err))
	}
	if token.Method != jwt.SigningMethodHS256 {
		return "", errors.New(fmt.Sprint("Invalid JWT signing method"))
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New(fmt.Sprint("Invalid JWT token"))
	}
	if claims["iss"] != "squrl" {
		return "", errors.New(fmt.Sprint("Invalid issuer ", claims["iss"]))
	}
	return claims["user"].(string), nil
}
