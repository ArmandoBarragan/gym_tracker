package controllers

import (
	"log"

	"github.com/ArmandoBarragan/gym_tracker/src/schemas"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "",
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := schemas.CreateUser{
		Username:             c.FormValue("username"),
		Password:             c.FormValue("password"),
		PasswordConfirmation: c.FormValue("passwordConfirmation"),
	}

	if user.Password != user.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Your passwords don't match",
		})
	}

	if !user.PasswordIsValid() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Password requires at least one special character and one digit",
		})
	}

	err := user.Create()

	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User create successfuly",
	})
}
