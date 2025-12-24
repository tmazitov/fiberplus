package behavior

import (
	"github.com/gofiber/fiber/v2"
	mods "github.com/tmazitov/fiberplus/internal/mods"
)

type CoreHandler[Services any, Input any, Output any] struct {
	DefaultHandler[Services]
}

func (h *CoreHandler[Services, Input, Output]) RequestBody(ctx *fiber.Ctx) (Input, bool) {
	value, ok := ctx.Locals("Input").(Input)
	return value, ok
}

func (h *CoreHandler[Services, Input, Output]) Reply(ctx *fiber.Ctx, response Output) {
	ctx.Locals("Output", response)
}

func (h *CoreHandler[Services, Input, Output]) Init() {
	h.mods = []mods.Mod{
		&mods.JsonInputMod[Input]{},
		&mods.JsonOutputMod[Output]{},
	}
}
