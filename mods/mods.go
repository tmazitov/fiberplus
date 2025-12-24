package mods

import "github.com/gofiber/fiber/v2"

type Mod interface {
	Setup() fiber.Handler
}
