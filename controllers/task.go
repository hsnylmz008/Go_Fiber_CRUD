package controllers

import (
 "context"

 "github.com/gofiber/fiber/v2"
 "github.com/your-username/your-project-name/models"
)

type TaskController struct {
 taskModel *models.TaskModel
}

func NewTaskController(taskModel *models.TaskModel) *TaskController {
 return &TaskController{taskModel: taskModel}
}

func (c *TaskController) GetTasks(ctx *fiber.Ctx) error {
 tasks, err := c.taskModel.GetAllTasks(ctx.Context())
 if err != nil {
  return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "Error fetching tasks",
  })
 }

 return ctx.JSON(tasks)
}

func (c *TaskController) GetTask(ctx *fiber.Ctx) error {
 id, err := ctx.ParamsInt("id")
 if err != nil {
  return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "Invalid task ID",
  })
 }

 task, err := c.taskModel.GetTaskByID(ctx.Context(), id)
 if err != nil {
  if err.Error() == "no rows in result set" {
   return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
    "message": "Task not found",
   })
  }
  return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "Error fetching task",
  })
 }

 return ctx.JSON(task)
}

func (c *TaskController) CreateTask(ctx *fiber.Ctx) error {
 task := new(models.Task)
 if err := ctx.BodyParser(task); err != nil {
  return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "Invalid request payload",
  })
 }

 id, err := c.taskModel.CreateTask(ctx.Context(), task)
 if err != nil {
  return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "Error creating task",
  })
 }

 task.ID = id
 return ctx.JSON(task)
}

func (c *TaskController) UpdateTask(ctx *fiber.Ctx) error {
 id, err := ctx.ParamsInt("id")
 if err != nil {
  return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "Invalid task ID",
  })
 }

 task := new(models.Task)
 if err := ctx.BodyParser(task); err != nil {
  return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "Invalid request payload",
  })
 }

 task.ID = id
 err = c.taskModel.UpdateTask(ctx.Context(), task)
 if err != nil {
  return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "Error updating task",
  })
 }

 return ctx.JSON(task)
}

func (c *TaskController) DeleteTask(ctx *fiber.Ctx) error {
 id, err := ctx.ParamsInt("id")
 if err != nil {
  return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
   "message": "Invalid task ID",
  })
 }

 err = c.taskModel.DeleteTask(ctx.Context(), id)
 if err != nil {
  return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
   "message": "Error deleting task",
  })
 }

 return ctx.SendStatus(fiber.StatusNoContent)
}
