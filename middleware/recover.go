package middleware

import (
	response "auth/backend/utils"
	"fmt"
	"log"
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
)

// Middleware to handle "panics" and return custom responses
func CustomRecoverMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			panicMessage := fmt.Sprintf("%v", r)
			log.Printf("Recovered from panic: %v", panicMessage)

			// trace detail error go runtime
			stackTrace := debug.Stack()
			log.Printf("Panic problem: %v", string(stackTrace))
			c.Status(500).JSON(response.ResponseMeta{
				StatusCode: 500,
				Message:    "Error panic",
				Error:      panicMessage,
				Data:       string(stackTrace),
			})
		}
	}()
	return c.Next()
}
