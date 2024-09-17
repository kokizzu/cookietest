package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

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
		strVal := c.Cookies("test", "0")
		intVal, _ := strconv.Atoi(strVal)
		c.Cookie(&fiber.Cookie{
			Name:     "test",
			Value:    strconv.Itoa(intVal + 1),
			Path:     "/",
			Domain:   ".local.benalu.dev",
			Secure:   true,
			HTTPOnly: true,
			SameSite: "Lax",
		})
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":13001"))
}
