package main

import (
	"encoding/json"

	"ambassador/src/database"
	"ambassador/src/events"
	"ambassador/src/routes"
	"ambassador/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	database.SetupRedis()
	database.SetupCacheChannel()
	services.Setup()
	events.SetupProducer()

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
