package handler

import (
	"auth/backend/middleware"
	"auth/backend/model"
	"auth/backend/services"
	"auth/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandlerMethod struct {
	service *services.AuthServiceMethod
	log     *logrus.Logger
}

func AuthHandler(service *services.AuthServiceMethod, log *logrus.Logger) *AuthHandlerMethod {
	return &AuthHandlerMethod{service, log}
}

func (h *AuthHandlerMethod) LoginHdlr(c *fiber.Ctx) error {
	h.log.Println("Execute function LoginHdlr")

	var payloadLogin model.Login
	if err := c.BodyParser(&payloadLogin); err != nil {
		h.log.Println("Failed parsed payload data login")
		return c.Status(400).JSON(utils.ResponseData{
			StatusCode: 400,
			Message:    "Error parsed payload data login",
			Error:      err.Error(),
		})
	}

	// validate data
	errorFields, validationError := middleware.ValidateData(&payloadLogin)
	if validationError != nil {
		h.log.Println("Validation error in LoginHdlr:", validationError)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
	}

	loginData, err := h.service.LoginSvc(payloadLogin)
	if err != nil {
		h.log.Println("Failed save data in LoginHdlr")
		return c.Status(404).JSON(utils.ResponseData{
			StatusCode: 404,
			Message:    "Error login",
			Error:      err.Error(),
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully Login!",
		Data:       loginData,
	})

}

func (h *AuthHandlerMethod) VerifyOTPHdlr(c *fiber.Ctx) error {
	h.log.Println("Execute function VerifyOTPHdlr")

	var payloadVerify model.VerifyOTP
	if err := c.BodyParser(&payloadVerify); err != nil {
		h.log.Println("Failed parsed payload data verify OTP")
		return c.Status(400).JSON(utils.ResponseData{
			StatusCode: 400,
			Message:    "Error parsed payload data verify OTP",
			Error:      err.Error(),
		})
	}

	// validate data
	errorFields, validationError := middleware.ValidateData(&payloadVerify)
	if validationError != nil {
		h.log.Println("Validation error in VerifyOTPHdlr:", validationError)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
	}

	verifyOTP, err := h.service.VerifyOTPSvc(payloadVerify)
	if err != nil {
		h.log.Println("Failed save data in VerifyOTPHdlr")
		return c.Status(404).JSON(utils.ResponseData{
			StatusCode: 404,
			Message:    "Error verify OTP",
			Error:      err.Error(),
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully verify OTP",
		Data:       verifyOTP,
	})
}
