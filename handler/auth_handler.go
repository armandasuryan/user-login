package handler

import (
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
	errorFields, validationError := utils.ValidateData(&payloadLogin)
	if validationError != nil {
		h.log.Println("Validation error in LoginHdlr:", validationError)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
	}

	loginData, err := h.service.LoginSvc(payloadLogin)
	if err != "" {
		h.log.Println("Failed get data in LoginHdlr")
		return c.Status(401).JSON(utils.ResponseData{
			StatusCode: 401,
			Message:    err,
			Error:      "Error login",
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
	errorFields, validationError := utils.ValidateData(&payloadVerify)
	if validationError != nil {
		h.log.Println("Validation error in VerifyOTPHdlr:", validationError)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
	}

	verifyOTP, err := h.service.VerifyOTPSvc(payloadVerify)
	if err != "" {
		h.log.Println("Failed verify data in VerifyOTPHdlr")
		return c.Status(401).JSON(utils.ResponseData{
			StatusCode: 401,
			Message:    err,
			Error:      "Error verify OTP",
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully verify OTP",
		Data:       verifyOTP,
	})
}

func (h *AuthHandlerMethod) GetUserProfileHdlr(c *fiber.Ctx) error {
	h.log.Println("Execute function GetUserProfileHdlr")

	getUserProfile, err := h.service.GetUserProfileSvc(c)
	if err != nil {
		h.log.Println("Failed get user profile")
		return c.Status(401).JSON(utils.ResponseData{
			StatusCode: 401,
			Message:    "Failed get user profile",
			Error:      err.Error(),
		})
	}

	return c.Status(200).JSON(utils.ResponseData{
		StatusCode: 200,
		Message:    "Successfully get user profile",
		Data:       getUserProfile,
	})
}
