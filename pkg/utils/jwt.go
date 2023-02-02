package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userid uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userid,
	})

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

// returns user_id that is stored in the payload
func GetIdFromToken(token string) (uint, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil || !parsed.Valid {
		// token not valid
		return 0, errors.New("invalid token")
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)

	if !ok {
		// fail to extract claims
		return 0, errors.New("failed to extract claims from payload")
	}

	return uint(claims["user_id"].(float64)), nil
}
