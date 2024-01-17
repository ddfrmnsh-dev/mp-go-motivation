package main

import (
	"github.com/ddfrmnsh-dev/mp-go-motivation/routes"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils/migrations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	utils.DatabaseInit()

	migrations.Migration()

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		"message": "Hello",
	// 	})
	// })
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"message": "Hello",
	// 	})
	// })

	routes.RouteInit(app)

	app.Listen(":8000")

}
