package routes

import (
	c "github.com/Seifbarouni/go-service-registry/controllers"

	"github.com/gofiber/fiber/v2"
)

func GenerateRoutes(app *fiber.App) {
	app.Get("/", c.Index)
	app.Get("/services/:serviceName", c.GetService)
	app.Post("/services/:serviceName", c.AddService)
	app.Delete("/services/:serviceName", c.ServiceDown)
	app.Put("/services/:serviceName", c.ServiceUp)
}
