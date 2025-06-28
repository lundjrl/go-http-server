package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lundjrl/go-http-server/internal"
	"github.com/lundjrl/go-http-server/internal/model"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	internal.InitDatabaseConnection()

	app.Get("/api/v1/cat", model.GetCats)
	app.Get("/api/v1/cat/:id", model.GetCatById)
	app.Post("/api/v1/cat", model.CreateCat)
	app.Delete("/api/v1/cat/:id", model.DeleteCat)

	app.Get("/health", func(c *fiber.Ctx) error {
		msg := fiber.Map{"message": "still here!"}
		return c.Status(200).JSON(msg)
	})

	log.Fatal(app.Listen(":6660"))
}
