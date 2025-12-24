package app

import (
	"github.com/gofiber/fiber/v2"
)

type Handler[Services any] interface {
	Init()
	Handle(app *App[Services]) fiber.Handler
	IntroMods() []fiber.Handler
	OutroMods() []fiber.Handler
}
