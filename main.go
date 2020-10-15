package main

import (
	"log"
	"todoapi/database"
	"todoapi/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	app := fiber.New()

	database.Connect()

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":9876"))

	defer database.DB.Close()
}
