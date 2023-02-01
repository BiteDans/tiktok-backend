package utils

import (
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

// returns user_id in the payload, 0 if error
func ValidateJWT(token string) uint {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil || !parsed.Valid {
		// token not valid
		return 0
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)

	if !ok {
		// fail to extract claims
		return 0
	}

	return uint(claims["user_id"].(float64))
}
