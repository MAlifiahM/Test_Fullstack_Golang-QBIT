package comments

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CommentHandler struct {
	service *CommentService
}

func NewCommentHandler(service *CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) GetComments(c *fiber.Ctx) error {
	comments, err := h.service.GetComments()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error getting comments")
	}
	return c.JSON(comments)
}

func (h *CommentHandler) GetTotalComments(c *fiber.Ctx) error {
	total, err := h.service.CountTotalComments()
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error getting comments count")
	}
	return c.JSON(fiber.Map{
		"total_comments": total,
	})
}
