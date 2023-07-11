package main

import (
	"fmt"

	"github.com/goBackendFiber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// fmt.

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {

		fmt.Println(database.Client)

		return c.SendString("i am  response ")

	})
	app.Listen(":3000")

}
