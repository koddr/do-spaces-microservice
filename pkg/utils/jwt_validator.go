package utils

import (
	"fmt"
	"time"

	"github.com/koddr/do-spaces-microservice/pkg/repository"

	"github.com/gofiber/fiber/v2"
)

// TokenValidateExpireTime func for validating given JWT token with expire time.
func TokenValidateExpireTime(ctx *fiber.Ctx) (*TokenMetaData, error) {
	// Get claims from JWT.
	claims, errExtractTokenMetaData := ExtractTokenMetaData(ctx)
	if errExtractTokenMetaData != nil {
		// Return JWT parse error.
		return nil, errExtractTokenMetaData
	}

	// Checking, if now time greather than expiration from JWT.
	if time.Now().Unix() > claims.Expires {
		// Return unauthorized (permission denied) error message.
		return nil, fmt.Errorf(repository.UnauthorizedExpirationTime)
	}

	return claims, nil
}

// TokenValidateExpireTimeAndCredentials func for validating given JWT token with expire time and credentials.
func TokenValidateExpireTimeAndCredentials(ctx *fiber.Ctx, credentials []string) (*TokenMetaData, error) {
	// Get claims from JWT.
	claims, errExtractTokenMetaData := ExtractTokenMetaData(ctx)
	if errExtractTokenMetaData != nil {
		// Return JWT parse error.
		return nil, errExtractTokenMetaData
	}

	// Checking, if now time greather than expiration from JWT.
	if time.Now().Unix() > claims.Expires {
		// Return unauthorized (permission denied) error message.
		return nil, fmt.Errorf(repository.UnauthorizedExpirationTime)
	}

	//
	for _, credential := range credentials {
		// Return unauthorized (permission denied) error message.
		if !SearchStringInArray(credential, claims.Credentials) {
			return nil, fmt.Errorf(repository.UnauthorizedCredentials)
		}
	}

	return claims, nil
}
