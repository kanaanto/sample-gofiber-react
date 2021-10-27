package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kanaanto/sample-gofiber-react/routes"
)

func setupRoutes(app *fiber.App) {
	// give response when at /
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "😉",
		})
	})

	// api group
	api := app.Group("/api")

	// give response when at /api
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "😉",
		})
	})

	// connect todo routes
	routes.Route(api.Group("/transactions"))
}

func main() {
	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	// Listen on server 8000 and catch error if any
	err := app.Listen(":8000")

	// handle error
	if err != nil {
		panic(err)
	}
}
