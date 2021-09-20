package main

import (
	"os"

	"github.com/koddr/do-spaces-microservice/pkg/configs"
	"github.com/koddr/do-spaces-microservice/pkg/middleware"
	"github.com/koddr/do-spaces-microservice/pkg/routes"
	"github.com/koddr/do-spaces-microservice/pkg/utils"

	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PrivateRoutes(app) // Register a private routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
