package middleware

import (
	"auth/backend/model"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(data model.ResponseLogin) (string, error) {
	fmt.Println("Execute function CreateJwt")

	claims := jwt.MapClaims{
		"id":        data.ID,
		"username":  data.Username,
		"email":     data.Email,
		"role_name": data.RoleName,
		"exp":       time.Now().Add(time.Hour * 3).Unix(), // exp in 3 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}
