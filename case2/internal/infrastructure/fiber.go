package infrastructure

import (
	"case2/internal/comments"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, commentHandler *comments.CommentHandler) {
	app.Get("/comments", commentHandler.GetComments)
	app.Get("/comments/total", commentHandler.GetTotalComments)
}
