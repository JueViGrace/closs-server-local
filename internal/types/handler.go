package types

import "github.com/gofiber/fiber/v2"

type AuthDataHandler = func(*fiber.Ctx, *AuthData) error
