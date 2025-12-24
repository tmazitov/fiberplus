package mods

import "github.com/gofiber/fiber/v2"

type ModeType int

const (
	IntroType = iota
	OutroType
)

type Mod interface {
	Setup() fiber.Handler
	Type() ModeType
}
