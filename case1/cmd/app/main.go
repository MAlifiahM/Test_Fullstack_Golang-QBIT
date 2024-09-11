package main

import (
	"case1/internal/fruits"
	"case1/internal/infrastructure"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	repo := fruits.NewJSONFruitRepository("fruits.json")
	service := fruits.NewFruitService(repo)
	handler := fruits.NewFruitHandler(service)

	infrastructure.SetupRouter(app, handler)
	
	app.Listen(":3000")
}
