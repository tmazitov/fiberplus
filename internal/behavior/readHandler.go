package behavior

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/internal/mods"
)

type ReadHandler[Services any, Input any] struct {
	DefaultHandler[Services]
}

func (h *ReadHandler[Services, Input]) Init() {
	h.mods = []mods.Mod{
		&mods.JsonInputMod[Input]{},
	}
}

func (h *ReadHandler[Services, Input]) RequestBody(ctx *fiber.Ctx) (Input, bool) {
	value, ok := ctx.Locals("Input").(Input)
	return value, ok
}
