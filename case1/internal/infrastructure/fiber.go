package infrastructure

import (
	"case1/internal/fruits"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App, fruitHandler *fruits.FruitHandler) {
	app.Get("/case1/answer1", fruitHandler.GetFruits)
	app.Get("/case1/answer23", fruitHandler.GetFruitsByType)
	app.Get("/case1/answer4", fruitHandler.GetUniqueFruits)
}
