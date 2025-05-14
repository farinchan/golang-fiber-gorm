package main

import (
	"golang-fiber-gorm/config"
	"golang-fiber-gorm/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "/public")

	config.ConnectDB()
	config.RunMigration()

	routes.AuthRoutes(app)
	routes.UserRoute(app)

	app.Listen(":8080")
}
