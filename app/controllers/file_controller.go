package controllers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/koddr/do-spaces-microservice/app/models"
	"github.com/koddr/do-spaces-microservice/pkg/utils"
	"github.com/koddr/do-spaces-microservice/platform/cdn"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

// GetObjectsListFromCDN method for return list of objects from CDN.
func GetObjectsListFromCDN(c *fiber.Ctx) error {
	// Get claims from JWT.
	_, err := utils.TokenValidateExpireTime(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create CDN connection.
	connDOSpaces, err := cdn.DOSpacesConnection()
	if err != nil {
		// Return status 500 and CDN connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create context with cancel.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // auto close

	// Get list of files from CDN.
	listObjectsChannel := connDOSpaces.ListObjects(
		ctx,
		os.Getenv("DO_SPACES_BUCKET_NAME"),
		minio.ListObjectsOptions{
			Prefix:    os.Getenv("DO_SPACES_UPLOADS_FOLDER_NAME"),
			Recursive: true,
		},
	)

	// Define File struct for object list.
	objects := []*models.FileFromCDN{}

	// Range object list from CDN for create a new Go object for JSON serialization.
	for object := range listObjectsChannel {
		// Check, if received object is valid.
		if object.Err != nil {
			// Return status 400 and bad request error.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   object.Err,
			})
		}

		// Skip upload folder from list, only files.
		if !strings.HasSuffix(object.Key, "/") {
			// Create a new File struct from object info.
			file := &models.FileFromCDN{
				Key:       object.Key,
				ETag:      object.ETag,
				VersionID: object.VersionID,
				URL:       fmt.Sprintf("%v/%v", os.Getenv("CDN_PUBLIC_URL"), object.Key),
			}

			// Add this file to objects list.
			objects = append(objects, file)
		}
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error":   false,
		"msg":     nil,
		"count":   len(objects),
		"objects": objects,
	})
}

// PutObjectToCDN method for upload an object with a specific type to CDN.
func PutObjectToCDN(c *fiber.Ctx) error {
	// Get claims from JWT.
	claims, err := utils.TokenValidateExpireTime(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create new LocalFile struct.
	localFile := &models.LocalFile{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(localFile); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new DO Spaces connection.
	connDOSpaces, err := cdn.DOSpacesConnection()
	if err != nil {
		// Return status 500 and CDN connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Define user ID.
	userID := claims.UserID.String()

	// Upload image file process.
	uploadObjectInfo, err := cdn.UploadObjectToCDN(connDOSpaces, localFile.Path, localFile.Type, userID)
	if err != nil {
		// Return status 400 and bad request error.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201 created.
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"info": fiber.Map{
			"key":        uploadObjectInfo.Key,
			"etag":       uploadObjectInfo.ETag,
			"size":       uploadObjectInfo.Size,
			"version_id": uploadObjectInfo.VersionID,
		},
		"url": fmt.Sprintf("%v/%v", os.Getenv("CDN_PUBLIC_URL"), uploadObjectInfo.Key),
	})
}

// RemoveObjectFromCDN method for remove exists object from CDN.
func RemoveObjectFromCDN(c *fiber.Ctx) error {
	// Get claims from JWT.
	_, err := utils.TokenValidateExpireTime(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create new FileFromCDN struct
	fileToDelete := &models.FileFromCDN{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(fileToDelete); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	//
	connDOSpaces, err := cdn.DOSpacesConnection()
	if err != nil {
		// Return status 500 and CDN connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Remove object from CDN.
	if errRemoveObject := connDOSpaces.RemoveObject(
		context.Background(),
		os.Getenv("DO_SPACES_BUCKET_NAME"),
		fileToDelete.Key,
		minio.RemoveObjectOptions{
			VersionID: fileToDelete.VersionID,
		},
	); errRemoveObject != nil {
		// Return status 400 and bad request error.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   errRemoveObject.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
