package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	fmt.Println("Started")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World  v4👋!")
	})

	app.Get("/late", func(c *fiber.Ctx) error {
		time.Sleep(2 * time.Second)
		return c.SendString("Hello late World  v4👋!")
	})

	app.Get("/password", func(c *fiber.Ctx) error {
		time.Sleep(2 * time.Second)
		return c.SendString("Your password is   v4👋!")
	})
	app.Listen(":80")

}
