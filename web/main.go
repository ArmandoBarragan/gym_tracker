package main

import (
	"fmt"

	"github.com/ArmandoBarragan/gym_tracker/conf"
	"github.com/ArmandoBarragan/gym_tracker/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var settings *conf.Settings = conf.InitSettings()
	conf.InitDatabase(settings.DB_DSN)
	app := fiber.New()
	routes.Auth(app)
	app.Listen(fmt.Sprintf(":%s", settings.PORT))
}
