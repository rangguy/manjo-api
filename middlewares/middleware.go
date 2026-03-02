package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

func responseUnauthorized(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  "error",
		"message": message,
	})
}

func validateAPIKey(c *fiber.Ctx) error {
	apiKey := c.Get("api-key")

	if apiKey == "" {
		return fmt.Errorf("missing required headers")
	}

	signatureKey := os.Getenv("SIGNATURE_KEY")

	hash := sha256.New()
	hash.Write([]byte(signatureKey))
	resultHash := hex.EncodeToString(hash.Sum(nil))

	if apiKey != resultHash {
		return fmt.Errorf("invalid api key")
	}

	return nil
}

func Authenticate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := validateAPIKey(c); err != nil {
			return responseUnauthorized(c, err.Error())
		}

		return c.Next()
	}
}
