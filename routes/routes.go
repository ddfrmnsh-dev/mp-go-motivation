package routes

import (
	"github.com/ddfrmnsh-dev/mp-go-motivation/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Get("/motivation", controllers.GetAllMotivationsRandom)
	r.Get("/random", controllers.GetAllMotivationsRandomRender)
	r.Get("/", controllers.GetHtml)
	r.Get("/getAllMotivation", controllers.GetAllMotivation)
	r.Post("/create-motivation", controllers.AddNewQuote)
	// r.Post("/motivation", controllers.CreateMotivations)
	// r.Get("/motivation/:id", controllers.GetMotivationById)
	// r.Patch("/motivation/:id", controllers.UpdateMotivation)
	r.Delete("/motivation/:id", controllers.DeleteMotivation)
}
