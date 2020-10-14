package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", hello)

	log.Fatal(app.Listen(":3000"))

	//defer DB.Close()
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello Sergan")
}
