package router

import (
	"go_projects/controller/person"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	v1 := api.Group("/v1")
	_person := v1.Group("/person")
	_person.Post("find/name", person.GetUserLastName)
	_person.Post("create", person.CreatePerson)
	//_person.Post("/delete/:id", person.DeletePerson)
	//_person.Post("/update/:id", person.UpdatePerson)

	v1.Get("/getUserLastName/:name", person.GetUserLastName) // api/v1/getUserLastName/:name
}
