package routes

import (
	"user-service/backend/handler"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	App         *fiber.App
	UserHandler *handler.UserHandlerMethod
}

func (r *UserRoute) SetupUserRoute() {
	UserRoute := r.App.Group("/wms/api/v1")

	UserRoute.Post("/login", r.UserHandler.LoginUserHdlr)
	UserRoute.Post("login/verify-otp", r.UserHandler.VerifyOTPHdlr)
}
