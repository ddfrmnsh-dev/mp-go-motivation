package controllers

import (
	"math/rand"
	// "strconv"

	// "github.com/ddfrmnsh-dev/mp-go-motivation/models"
	// "github.com/ddfrmnsh-dev/mp-go-motivation/utils"
	// "github.com/go-playground/validator/v10"
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
