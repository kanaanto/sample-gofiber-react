package routes

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/kanaanto/sample-gofiber-react/controllers"
)

func Route(route fiber.Router) {
	route.Get("", controllers.GetTransactions)
}
