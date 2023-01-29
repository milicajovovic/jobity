package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(id int, email string) (string, error) {
	claims := jwt.MapClaims{
		"userId": strconv.Itoa(id),
		"email":  email,
		"exp":    time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}

func ParseToken(tokenString string) (string, error) {
	if token, err := ValidateToken(tokenString); err != nil {
		return "", err
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); !ok {
			return "", errors.New("unauthorized")
		} else {
			if token.Valid {
				return claims["userId"].(string), nil
			} else {
				return "", errors.New("unauthorized")
			}
		}
	}
}
