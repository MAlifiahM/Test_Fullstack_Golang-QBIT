package auth

import (
	"case3/internal/domain"
	"case3/internal/middleware/validation"
	"case3/internal/utilities"
	"github.com/gofiber/fiber/v2"
)

type HttpAuthHandler struct {
	authService domain.AuthService
}

func NewHttpHandler(r fiber.Router, authService domain.AuthService) {
	handler := &HttpAuthHandler{authService}
	r.Post("/login", validation.New[domain.ReqLogin](), handler.Login)
	r.Post("/register", validation.New[domain.ReqRegister](), handler.Register)
}

func (s *HttpAuthHandler) Login(c *fiber.Ctx) error {
	reqLogin := utilities.ExtractStructFromValidator[domain.ReqLogin](c)

	response, err := s.authService.Login(reqLogin.Email, reqLogin.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "login success",
		Error:   false,
		Data:    response,
	})
}

func (s *HttpAuthHandler) Register(c *fiber.Ctx) error {
	reqRegister := utilities.ExtractStructFromValidator[domain.ReqRegister](c)

	err := s.authService.Register(reqRegister.Username, reqRegister.Email, reqRegister.Password)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "register success",
		Error:   false,
	})
}
