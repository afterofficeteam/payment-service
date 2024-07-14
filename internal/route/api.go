package route

import (
	paymentHandler "payment-service/internal/module/payment/handler/rest"

	"payment-service/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// add /api prefix to all routes
	api := app.Group("/api")
	paymentHandler.NewHandler().Register(api)

	// health check route
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(response.Success(nil, "Server is running."))
	})

	// fallback route
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(response.Error("Route not found."))
	})
}
