package routes

import (
	"github.com/koddr/do-spaces-microservice/app/controllers"
	"github.com/koddr/do-spaces-microservice/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/v1", middleware.JWTProtected())

	// Routes for GET method:
	route.Get("/list", controllers.GetObjectsListFromCDN) // get file list from CDN

	// Routes for PUT method:
	route.Put("/upload", controllers.PutObjectToCDN) // upload object to CDN

	// Routes for DELETE method:
	route.Delete("/remove", controllers.RemoveObjectFromCDN) // remove object from CDN
}
