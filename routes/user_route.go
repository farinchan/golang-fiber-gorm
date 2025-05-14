package routes

import (
	"golang-fiber-gorm/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/users", controllers.CreateUser)
}
