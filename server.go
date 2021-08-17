package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewHttp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "welcome!",
		})
	})

	students := app.Group("/students", Auth())
	students.Get("/",
		Permission("students", PERMISSION_READ),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "all students",
				"data":    []struct{}{},
			})
		},
	)

	students.Post("/",
		Permission("students", PERMISSION_CREATE),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "student created",
				"data":    struct{}{},
			})
		},
	)

	students.Get("/:id",
		Permission("students", PERMISSION_READ),
		Permission("students", PERMISSION_CREATE),
		Permission("students", PERMISSION_UPDATE),
		Permission("students", PERMISSION_DELETE),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": fmt.Sprintf("student with id %s", c.Params("id", "0")),
				"data":    struct{}{},
			})
		},
	)

	return app
}
