package middleware

import (
	"auth/backend/model"
	"auth/backend/utils"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(data model.ResponseLogin) (string, error) {
	fmt.Println("Execute function CreateJwt")

	createJWTMap := jwt.MapClaims{
		"id":        data.ID,
		"username":  data.Username,
		"email":     data.Email,
		"role_name": data.RoleName,
		"exp":       time.Now().Add(time.Hour * 3).Unix(), // exp in 3 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, createJWTMap)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func JWTMiddleware(c *fiber.Ctx) error {

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(401).JSON(utils.ErrorResponse{
			StatusCode: 401,
			Message:    "Authorization header not provided",
			Error:      "Error get header authorization",
		})
	}

	jwtData, err := VerifyJWTToken(tokenString)
	if err != "" {
		return c.Status(401).JSON(utils.ErrorResponse{
			StatusCode: 401,
			Message:    err,
			Error:      "error verify token",
		})
	}

	// save data to local using context
	c.Locals("token", jwtData)

	// continue action
	return c.Next()
}

func VerifyJWTToken(tokenString string) (jwt.MapClaims, string) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, "Invalid or expired token"
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, ""
	} else {
		return nil, "Invalid token"
	}
}
