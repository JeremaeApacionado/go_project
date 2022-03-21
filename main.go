package main

import (
	"go_projects/router"

	"github.com/gofiber/fiber/v2"
)

var app *fiber.App

func main() {

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen((":3000"))
}
