package routes

import (
	"github.com/ArmandoBarragan/gym_tracker/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Auth(app *fiber.App) {
	route := app.Group("api/")
	route.Post("create-user", controllers.CreateUser)
	route.Post("login", controllers.Login)
}
