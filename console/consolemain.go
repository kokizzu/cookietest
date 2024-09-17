package main

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

//go:embed main.html
var mainHtml []byte

func main() {
	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.Send(mainHtml)
	})

	log.Fatal(app.Listen(":13000"))
}
