package routes

import (
 "github.com/gofiber/fiber/v2"
 "github.com/hsnylmz008/Go_Fiber_CRUD/controllers"
)

func SetupRoutes(app *fiber.App, taskController *controllers.TaskController) {
 app.Get("/tasks", taskController.GetTasks)
 app.Get("/tasks/:id", taskController.GetTask)
 app.Post("/tasks", taskController.CreateTask)
 app.Put("/tasks/:id", taskController.UpdateTask)
 app.Delete("/tasks/:id", taskController.DeleteTask)
}
