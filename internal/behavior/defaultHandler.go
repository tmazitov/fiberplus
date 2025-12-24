package behavior

import (
	"github.com/gofiber/fiber/v2"
	mods "github.com/tmazitov/fiberplus/internal/mods"
)

type DefaultHandler[Services any] struct {
	mods []mods.Mod
}

func (h *DefaultHandler[Services]) Init() {}

func (h *DefaultHandler[Services]) Mods() []fiber.Handler {

	if len(h.mods) == 0 {
		return nil
	}

	var handleFuncs = make([]fiber.Handler, 0, len(h.mods))

	for _, mod := range h.mods {
		handleFuncs = append(handleFuncs, mod.Setup())
	}

	return handleFuncs
}
