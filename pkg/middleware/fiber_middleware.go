package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	// Add middlewares.
	a.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins:     os.Getenv("ALLOW_ORIGINS"),
			AllowCredentials: true,
		}),
		// Add simple logger.
		logger.New(),
	)
}
