package behavior

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/tmazitov/fiberplus/internal/app"
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

func (h *ReadHandler[Services, Input]) Handle(app *app.App[Services]) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		log.Info("read:")

		return nil
	}
}
