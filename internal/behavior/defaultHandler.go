package behavior

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberplus "github.com/tmazitov/fiberplus/internal"
	mods "github.com/tmazitov/fiberplus/internal/mods"
)

type DefaultHandler[Services any] struct {
	introMods []mods.Mod
	outroMods []mods.Mod
}

func (h *DefaultHandler[Services]) Init() {
	h.introMods = nil
	h.outroMods = nil
}

func (h *DefaultHandler[Services]) Handle(app *fiberplus.App[Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Info("default:")

		return nil
	}
}

func (h *DefaultHandler[Services]) IntroMods() []fiber.Handler {

	if len(h.introMods) == 0 {
		return nil
	}

	var handleFuncs = make([]fiber.Handler, 0, len(h.introMods))

	for _, mod := range h.introMods {
		handleFuncs = append(handleFuncs, mod.Setup())
	}

	return handleFuncs
}

func (h *DefaultHandler[Services]) OutroMods() []fiber.Handler {

	if len(h.outroMods) == 0 {
		return nil
	}

	var handleFuncs = make([]fiber.Handler, 0, len(h.outroMods))

	for _, mod := range h.outroMods {
		handleFuncs = append(handleFuncs, mod.Setup())
	}

	return handleFuncs
}
