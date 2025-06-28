package model

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lundjrl/go-http-server/internal/database"
	"gorm.io/gorm"
)

type Cat struct {
	gorm.Model
	Name       string `json:"name"`
	Age        int    `json:"age"`
	HairLength string `json:"hair_length"`
	HairColor  string `json:"hair_color"`
}

func GetCats(c *fiber.Ctx) error {
	db := database.DBConn
	var cats []Cat
	db.Find(&cats)
	return c.JSON(cats)
}

func GetCatById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var cat Cat
	db.Find(&cat, id)
	return c.JSON(cat)
}

func CreateCat(c *fiber.Ctx) error {
	db := database.DBConn
	cat := new(Cat)

	if err := c.BodyParser(cat); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&cat)
	return c.JSON(cat)
}

func DeleteCat(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var cat Cat
	db.First(&cat, id)

	if cat.Name == "" {
		return c.Status(500).SendString("No Cat Found with ID")
	}
	db.Delete(&cat)
	return c.SendString("Cat Successfully deleted")
}
