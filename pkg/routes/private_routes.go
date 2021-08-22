package routes

import (
	"Komentory/cdn/app/controllers"
	"Komentory/cdn/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:
	route.Get("/list", middleware.JWTProtected(), controllers.GetFileListFromCDN) // get file list from CDN

	// Routes for PUT method:
	route.Put("/upload/image", middleware.JWTProtected(), controllers.PutImageFileToCDN) // upload image file to CDN

	// Routes for DELETE method:
	route.Delete("/remove", middleware.JWTProtected(), controllers.RemoveFileFromCDN) // delete one file from CDN
}
