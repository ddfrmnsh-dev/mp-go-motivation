package controllers

import (
	"strconv"

	"math/rand"

	"github.com/ddfrmnsh-dev/mp-go-motivation/models"
	"github.com/ddfrmnsh-dev/mp-go-motivation/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllMotivationsRandom(c *fiber.Ctx) error {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	data := answers[rand.Intn(len(answers))]

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "ok",
		"motivation": data,
	})
}

func GetHtml(c *fiber.Ctx) error {
	return c.Render("index", nil)
}

func GetAllMotivation(c *fiber.Ctx) error {
	var motivation []*models.Motivation

	if err := utils.DB.Find(&motivation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendStatus(404)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "OK",
		"data":    motivation,
	})
}
func GetAllMotivationsRandomRender(c *fiber.Ctx) error {
	var motivation []*models.Motivation

	// utils.DB.Debug().Find(&motivation)

	if err := utils.DB.Find(&motivation).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	if len(motivation) == 0 {
		// Tangani situasi jika tidak ada data
		return c.Status(fiber.StatusNotFound).SendString("No data found")
	}

	data := motivation[rand.Intn(len(motivation))]
	// return c.Render("index", fiber.Map{
	// 	"data": data,
	// })

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "ok",
		"motivation": data,
	})
}

func AddNewQuote(c *fiber.Ctx) error {
	mov := new(models.Motivation)

	if err := c.BodyParser(mov); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(err.Error())
	}

	name := c.FormValue("name")
	quote := c.FormValue("quote")

	validate := validator.New()

	errValid := validate.Struct(mov)
	if errValid != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed to validate",
			"error":   errValid.Error(),
		})
	}

	if mov.Name == "" || mov.Quote == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name and email are required",
		})
	}

	newMov := models.Motivation{
		Name:  name,
		Quote: quote,
	}

	utils.DB.Debug().Create(&newMov)

	return c.Redirect("/")
	// return c.Status(fiber.StatusOK).JSON(fiber.Map{
	// 	"message": "success created new user",
	// })
}

func DeleteMotivation(c *fiber.Ctx) error {
	mov := new(models.Motivation)

	id, _ := strconv.Atoi(c.Params("id"))

	utils.DB.Debug().Where("id = ?", id).Delete(&mov)

	return c.Redirect("/")
}
