package middleware

import (
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
	"github.com/jkarlos000/sc-market/api/conf"
)

func Protected() func(*fiber.Ctx) {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusBadRequest)
		c.JSON(fiber.Map{"status": "error", "message": "Token invalido o incorrecto", "data": nil})

	} else {
		c.Status(fiber.StatusUnauthorized)
		c.JSON(fiber.Map{"status": "error", "message": "Token Expirado", "data": nil})
	}
}
