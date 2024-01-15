package main

import (
	"github.com/ddfrmnsh-dev/mp-go-motivation/routes"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils/migrations"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	utils.DatabaseInit()

	migrations.Migration()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "Hello",
		})
	})

	routes.RouteInit(app)

	app.Listen(":8000")

}
