package person

import (
	"go_projects/controller/person/model"
	"go_projects/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePerson(c *fiber.Ctx) error {
	payload := new(model.PersonModel)
	util.BodyParser(c, payload)

	fname := payload.FirstName
	lname := payload.LastName
	age := strconv.Itoa(payload.Age)

	query := "INSERT INTO person (first_name, last_name, age) VALUE(" + fname + ", " + lname + ", " + age + ")"

	//person, _ := json.Marshal(payload)

	return c.SendString(query)
}
