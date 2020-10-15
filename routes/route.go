package routes

import (
	"todoapi/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/items", handler.GetAllTasks)
	app.Post("/api/v1/item", handler.AddTask)
	app.Put("/api/v1/item/:id", handler.UpdateTask)
	app.Delete("/api/v1/item/:id", handler.DeleteTask)
}
