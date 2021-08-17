package main

import (
	"github.com/gofiber/fiber/v2"
)

func Auth() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fromHeader := c.Get("authorization")

		if fromHeader == "" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"errors": "access denied",
			})
		}

		c.Locals("user", fromHeader)
		return c.Next()
	}
}

func Permission(rules ...string) func(*fiber.Ctx) error {
	if len(rules) != 2 {
		panic("length of permission rule must be 2")
	}

	iam := NewIam()

	return func(c *fiber.Ctx) error {
		user := c.Locals("user")
		ok, err := iam.Enforce(user, rules[0], rules[1])

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"errors": "error occured while authorizing user",
			})
		}
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"errors": "access denied",
			})
		}

		return c.Next()
	}
}
