package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	msg := "GO Server listening on port 4000"

	log.Fatal(app.Listen(":4000"))
	fmt.Println(msg)
}
