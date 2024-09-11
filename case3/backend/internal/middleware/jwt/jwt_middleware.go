package jwt

import (
	"case3/internal/domain"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(secret)},
		ContextKey:   "user",
		ErrorHandler: jwtErrorHandler,
	})
}

func jwtErrorHandler(c *fiber.Ctx, err error) error {
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(domain.ResponseDefault{
			Code:    fiber.StatusUnauthorized,
			Message: "unauthorized",
			Error:   true,
		})
	}

	return nil
}
