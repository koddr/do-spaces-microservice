package utils

import (
	"fmt"
	"time"

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
	if time.Now().Unix() > claims.Expire {
		// Return unauthorized (permission denied) error message.
		return nil, fmt.Errorf("token was expired")
	}

	return claims, nil
}
