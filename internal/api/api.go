package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/bozkayasalihx/repartners-test-case/internal/handlers"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://159.65.241.100:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Static("/", "internal/web/static/index.html")

	app.Post("/calculate", handlers.HandlePacks)

	log.Fatal(app.Listen(":3000"))
}
