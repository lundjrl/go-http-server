package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		msg := fiber.Map{"message": "still here!"}
		return c.Status(200).JSON(msg)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		msg := fiber.Map{"message": "still here!"}
		return c.Status(200).JSON(msg)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		msg := fiber.Map{"message": "test"}
		return c.Status(200).JSON(msg)
	})

	log.Fatal(app.Listen(":6660")) // ENV in GO? Shared ENV for docker and here?
}
