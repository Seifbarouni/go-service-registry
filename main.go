package main

import (
	"log"

	"github.com/Seifbarouni/go-service-registry/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.GenerateRoutes(app)

	log.Fatal(app.Listen(":8671"))
}
