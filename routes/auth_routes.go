package routes

import (
	"golang-fiber-gorm/controllers"
	middleware "golang-fiber-gorm/midleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRoutes(app *fiber.App) {
    app.Post("/login", controllers.Login)

    app.Get("/profile", middleware.Protected(), func(c *fiber.Ctx) error {
        token := c.Locals("user").(*jwt.Token)
        claims := token.Claims.(jwt.MapClaims)
        email := claims["email"]
        return c.JSON(fiber.Map{"message": "Halo!", "email": email})
    })
}