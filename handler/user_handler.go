package handler

import (
	"user-service/backend/middleware"
	"user-service/backend/model"
	"user-service/backend/services"
	"user-service/backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandlerMethod struct {
	service *services.UserServiceMethod
	log     *logrus.Logger
}

func UserHandler(service *services.UserServiceMethod, log *logrus.Logger) *UserHandlerMethod {
	return &UserHandlerMethod{service, log}
}

func (h *UserHandlerMethod) LoginUserHdlr(c *fiber.Ctx) error {
	h.log.Println("Execute function LoginUserHdlr")

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
		h.log.Println("Validation error in LoginUserHdlr:", validationError)
		return c.Status(400).JSON(utils.ResponseValidator{
			StatusCode: 400,
			Message:    "Fill The Required Fields",
			Error:      errorFields,
		})
	}

	loginData, err := h.service.LoginUserSvc(payloadLogin)
	if err != nil {
		h.log.Println("Failed save data in LoginUserHdlr")
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

func (h *UserHandlerMethod) VerifyOTPHdlr(c *fiber.Ctx) error {
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
