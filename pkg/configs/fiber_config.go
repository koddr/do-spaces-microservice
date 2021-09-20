package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server read timeout.
	readTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		return fiber.Config{}
	}

	// Define server write timeout.
	writeTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if err != nil {
		return fiber.Config{}
	}

	// Define server idle timeout.
	idleTimeoutSecondsCount, err := strconv.Atoi(os.Getenv("SERVER_IDLE_TIMEOUT"))
	if err != nil {
		return fiber.Config{}
	}

	// Define server settings for production.
	var startupMessage bool
	if os.Getenv("STAGE_STATUS") == "prod" {
		startupMessage = true
	}

	// Return Fiber configuration.
	return fiber.Config{
		DisableStartupMessage: startupMessage,
		ReadTimeout:           time.Duration(readTimeoutSecondsCount) * time.Second,
		WriteTimeout:          time.Duration(writeTimeoutSecondsCount) * time.Second,
		IdleTimeout:           time.Duration(idleTimeoutSecondsCount) * time.Second,
	}
}
