package product

import (
	"case3/internal/domain"
	"case3/internal/middleware/validation"
	"case3/internal/utilities"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type HttpProductHandler struct {
	productService domain.ProductService
}

func NewHttpPublicProductHandler(r fiber.Router, productService domain.ProductService) {
	handler := &HttpProductHandler{productService}
	r.Get("/", handler.GetAll)
	r.Get("/:id", handler.GetByID)
}

func NewHttpPrivateProductHandler(r fiber.Router, productService domain.ProductService) {
	handler := &HttpProductHandler{productService}
	r.Post("/", validation.New[domain.RequestProduct](), handler.CreateProduct)
	r.Patch("/:id", validation.New[domain.RequestProduct](), handler.UpdateProduct)
	r.Delete("/:id", handler.DeleteProduct)
}

func (s *HttpProductHandler) GetAll(c *fiber.Ctx) error {
	products, err := s.productService.GetAll()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "success",
		Error:   false,
		Data:    products,
	})
}

func (s *HttpProductHandler) GetByID(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ResponseDefault{
			Code:    fiber.StatusBadRequest,
			Message: err.Error(),
			Error:   true,
		})
	}

	product, err := s.productService.GetByID(productId)

	if err != nil {
		if err.Error() == fmt.Sprintf("record not found") {
			return c.Status(fiber.StatusNotFound).JSON(domain.ResponseDefault{
				Code:    fiber.StatusNotFound,
				Message: err.Error(),
				Error:   true,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "success",
		Error:   false,
		Data:    product,
	})
}

func (s *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	reqProduct := utilities.ExtractStructFromValidator[domain.RequestProduct](c)

	err := s.productService.CreateProduct(reqProduct)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(domain.ResponseDefault{
		Code:    fiber.StatusCreated,
		Message: "create product success",
		Error:   false,
	})
}

func (s *HttpProductHandler) UpdateProduct(c *fiber.Ctx) error {
	productId, _ := c.ParamsInt("id")

	reqProduct := utilities.ExtractStructFromValidator[domain.RequestProduct](c)

	err := s.productService.UpdateProduct(productId, reqProduct)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "update product success",
		Error:   false,
	})
}

func (s *HttpProductHandler) DeleteProduct(c *fiber.Ctx) error {
	productId, _ := c.ParamsInt("id")

	err := s.productService.DeleteProduct(productId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(domain.ResponseDefault{
			Code:    fiber.StatusInternalServerError,
			Message: err.Error(),
			Error:   true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseDefault{
		Code:    fiber.StatusOK,
		Message: "delete product success",
		Error:   false,
	})
}
