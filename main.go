package main

import (
	"log"
	"os"

	"github.com/Seifbarouni/go-service-registry/database"
	"github.com/Seifbarouni/go-service-registry/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitializeDB(os.Getenv("DB_CONN"))

	engine := html.New("./views",".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	routes.GenerateRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
