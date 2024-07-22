package middleware

import (
	"rucara-services/dtos"
	utils "rucara-services/helpers"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Missing or malformed JWT")
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT")
			}
			return dtos.JWT_KEY, nil
		})

		if err != nil {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid JWT")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userID", claims["userID"])
		} else {
			return utils.ErrorResponse(c, fiber.StatusUnauthorized, "Invalid JWT")
		}

		return c.Next()
	}
}
