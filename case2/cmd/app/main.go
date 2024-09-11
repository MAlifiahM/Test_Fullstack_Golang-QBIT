package main

import (
	"case2/internal/comments"
	"case2/internal/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	repo := comments.NewJSONCommentRepository("comments.json")
	service := comments.NewCommentService(repo)
	handler := comments.NewCommentHandler(service)

	infrastructure.SetupRouter(app, handler)

	app.Listen(":3000")
}
