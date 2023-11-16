package utils

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUser struct {
	jwt.RegisteredClaims
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetAuthUser(c *fiber.Ctx) AuthUser {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	email := claims["email"].(string)
	idString := claims["id"].(string)
	id, _ := strconv.Atoi(idString)
	return AuthUser{
		ID:    uint64(id),
		Name:  name,
		Email: email,
	}
}
