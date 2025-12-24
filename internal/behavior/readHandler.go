package behavior

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberplus "github.com/tmazitov/fiberplus/internal"
	"github.com/tmazitov/fiberplus/internal/mods"
)

type ReadHandler[Services any, Input any] struct {
	DefaultHandler[Services]
}

func (h *ReadHandler[Services, Input]) Init() {
	h.introMods = []mods.Mod{
		&mods.JsonInputMod[Input]{},
	}
}

func (h *ReadHandler[Services, Input]) Handle(app *fiberplus.App[Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Info("read:")

		return nil
	}
}
