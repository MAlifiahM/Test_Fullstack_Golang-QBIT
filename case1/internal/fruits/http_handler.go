package fruits

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FruitHandler struct {
	service *FruitService
}

func NewFruitHandler(service *FruitService) *FruitHandler {
	return &FruitHandler{service: service}
}

func (h *FruitHandler) GetFruits(c *fiber.Ctx) error {
	fruits, err := h.service.GetFruits()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error getting fruits")
	}

	return c.JSON(fruits)
}

func (h *FruitHandler) GetFruitsByType(c *fiber.Ctx) error {
	fruits, err := h.service.GetFruitsByType()

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error getting fruits")
	}

	return c.JSON(fruits)
}

func (h *FruitHandler) GetUniqueFruits(c *fiber.Ctx) error {
	uniqueFruits, err := h.service.GetUniqueFruits()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error getting unique fruits")
	}

	return c.JSON(uniqueFruits)
}
