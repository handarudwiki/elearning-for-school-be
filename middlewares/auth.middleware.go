package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/handarudwiki/services"
)

func CheckAuth(jwtService services.JWTService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println(jwtService)
		authorization := ctx.Get("Authorization")

		if authorization == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		bearer := strings.Split(authorization, "Bearer ")

		if len(bearer) != 2 {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		token := bearer[1]

		claims, err := jwtService.ValidateToken(token)
		fmt.Println(claims)

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		userID := int(claims["userId"].(float64))
		role := int(claims["role"].(float64))

		ctx.Locals("userId", userID)
		ctx.Locals("role", role)

		return ctx.Next()

	}
}
