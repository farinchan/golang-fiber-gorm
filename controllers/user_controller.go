package controllers

import (
	"golang-fiber-gorm/config"
	"golang-fiber-gorm/models"

	"github.com/gofiber/fiber/v2"
)

var db = config.DB

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	db.Create(&user)
	return c.JSON(user)
}
