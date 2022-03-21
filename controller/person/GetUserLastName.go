package person

import (
	"github.com/gofiber/fiber/v2"
)

func GetUserLastName(c *fiber.Ctx) error {
	name := c.Params("name")

	if name == "Jeremae" {
		return c.JSON(Hello("Apacionado"))
	}
	return c.SendString(Hello("not found"))
}
