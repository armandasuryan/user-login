package routes

import (
	"auth/backend/handler"
	"auth/backend/middleware"

	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	App         *fiber.App
	AuthHandler *handler.AuthHandlerMethod
}

func (r *AuthRoute) SetupAuthRoute() {
	authRoute := r.App.Group("/wms/api/v1")

	authRoute.Post("/login", r.AuthHandler.LoginHdlr)
	authRoute.Post("/login/verify-otp", r.AuthHandler.VerifyOTPHdlr)
	authRoute.Get("/user-profile", middleware.JWTMiddleware, r.AuthHandler.GetUserProfileHdlr)
}
