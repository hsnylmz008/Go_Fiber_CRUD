package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/hsnylmz008/Go_Fiber_CRUD/models"
	"github.com/hsnylmz008/Go_Fiber_CRUD/controllers"
	"github.com/hsnylmz008/Go_Fiber_CRUD/routes"
)

func setupDatabase() (*pgxpool.Pool, error) {
	connString := "postgresql://myuser:mypassword@localhost:5432/mydatabase"
	pool, err := pgxpool.Connect(nil, connString)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func main() {
	db, err := setupDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	app := fiber.New()

	// Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Models
	taskModel := models.NewTaskModel(db)

	// Controllers
	taskController := controllers.NewTaskController(taskModel)

	// Routes
	routes.SetupRoutes(app, taskController)

	// Start server
	err = app.Listen(":3000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
