package cart

import (
	"case3/internal/domain"
	"case3/internal/middleware/validation"
	"case3/internal/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type HttpCartHandler struct {
	cartService domain.CartService
}

func NewHttpCartHandler(r fiber.Router, cartService domain.CartService) {
	handler := &HttpCartHandler{cartService}
	r.Get("/", handler.GetCartByUserID)
	r.Delete("/:id", handler.DeleteCartByUserID)
	r.Post("/", validation.New[domain.RequestCart](), handler.CreateCart)
	r.Patch("/:id", validation.New[domain.RequestCart](), handler.UpdateCart)
}

func (s *HttpCartHandler) GetCartByUserID(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	carts, err := s.cartService.GetCartByUserID(userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "get cart success",
		Error:   false,
		Data:    carts,
	})
}

func (s *HttpCartHandler) DeleteCartByUserID(c *fiber.Ctx) error {
	cartId, _ := c.ParamsInt("id")

	err := s.cartService.DeleteCartByUserID(cartId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "delete cart success",
		Error:   false,
	})
}

func (s *HttpCartHandler) CreateCart(c *fiber.Ctx) error {
	claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
	userId := int(claims["id"].(float64))

	reqCart := utilities.ExtractStructFromValidator[domain.RequestCart](c)

	err := s.cartService.CreateCart(reqCart, userId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(domain.ResponseDefault{
		Code:    fiber.StatusCreated,
		Message: "create cart success",
		Error:   false,
	})
}

func (s *HttpCartHandler) UpdateCart(c *fiber.Ctx) error {
	cartId, _ := c.ParamsInt("id")

	reqCart := utilities.ExtractStructFromValidator[domain.RequestCart](c)

	err := s.cartService.UpdateCart(cartId, reqCart)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "update cart success",
		Error:   false,
	})
}
