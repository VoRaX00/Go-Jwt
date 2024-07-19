package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"site-cv-back/database"
	"site-cv-back/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "POST, GET, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	err := app.Listen(":8000")
	if err != nil {
		return
	}
}
