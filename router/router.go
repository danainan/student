package router

import (
	"student/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())
	api.Get("/", handler.Hello)
	api.Get("/GetAllStudent", handler.GetAllStudent)
	api.Get("/GetById/:id", handler.GetById)
	api.Post("/CreateStudent", handler.CreateStudent)
	api.Put("/UpdateStudent/:id", handler.UpdateStudent)
	api.Delete("/DeleteStudent/:id", handler.DeleteStudent)
}
