package behavior

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tmazitov/fiberplus/internal/mods"
)

type WriteHandler[Services any, Output any] struct {
	DefaultHandler[Services]
}

func (h *WriteHandler[Services, Output]) Init() {
	h.mods = []mods.Mod{
		&mods.JsonOutputMod[Output]{},
	}
}

func (h *WriteHandler[Services, Output]) Reply(ctx *fiber.Ctx, response Output) {
	ctx.Locals("Output", response)
}
