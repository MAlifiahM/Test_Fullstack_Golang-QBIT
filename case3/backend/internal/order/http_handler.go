package order

import (
	"case3/internal/domain"
	"case3/internal/middleware/validation"
	"case3/internal/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type HttpOrderHandler struct {
	OrderService domain.OrderService
}

func NewHttpOrderHandler(r fiber.Router, orderService domain.OrderService) {
	handler := &HttpOrderHandler{orderService}
	r.Get("/", handler.GetOrderByUserID)
	r.Post("/", validation.New[domain.RequestOrder](), handler.CreateOrder)
}

func (s *HttpOrderHandler) GetOrderByUserID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	orders, err := s.OrderService.GetOrderByUserID(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "get order success",
		Error:   false,
		Data:    orders,
	})
}

func (s *HttpOrderHandler) CreateOrder(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	reqOrder := utilities.ExtractStructFromValidator[domain.RequestOrder](c)

	err := s.OrderService.CreateOrder(reqOrder, userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "create order success",
		Error:   false,
	})
}
