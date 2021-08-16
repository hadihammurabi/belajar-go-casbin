package main

import (
	"github.com/gofiber/fiber/v2"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fromHeader := c.Get("authorization")

		if fromHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "access denied",
			})
		}

		c.Locals("user", fromHeader)
		return c.Next()
	}
}

