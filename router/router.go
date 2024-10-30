package router

import (
	"student/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpRoutes(app *fiber.App) {
	api := app.Group("api", logger.New())

	studentHandler := studenthandler.NewStudentHandler()

	api.Get("/GetAllStudent", studentHandler.GetAllStudent)
	api.Get("/GetById/:student_id", studentHandler.GetById)
	api.Post("/CreateStudent", studentHandler.CreateStudent)
	api.Put("/UpdateStudent/:student_id", studentHandler.UpdateStudent)
	api.Delete("/DeleteStudent/:student_id", studentHandler.DeleteStudent)
}
