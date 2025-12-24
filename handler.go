package fiberplus

import (
	"github.com/gofiber/fiber/v2"
)

type Handler[Services any] interface {
	Init()
	Mods() []fiber.Handler
	Handle(app *App[Services]) fiber.Handler
}
