package main

import (
	_ "embed"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	corsRule := cors.ConfigDefault
	corsRule.AllowCredentials = true // required
	corsRule.AllowOrigins = "https://console.local.benalu.dev"
	app.Use(cors.New(corsRule))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/html")
		return c.Send(mainHtml)
	})

	log.Fatal(app.Listen(":13000"))
}
