package main

import (
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
		Permission("students", PERMISSION_CREATE),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "all students",
				"data":    []bool{},
			})
		},
	)
	students.Post("/",
		Permission("students", PERMISSION_CREATE),
		func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"message": "student created",
				"data":    []bool{},
			})
		},
	)

	return app
}
