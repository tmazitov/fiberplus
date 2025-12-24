package behavior

import "github.com/tmazitov/fiberplus/internal/mods"

type WriteHandler[Services any, Output any] struct {
	DefaultHandler[Services]
}

func (h *WriteHandler[Services, Output]) Init() {
	h.outroMods = []mods.Mod{
		&mods.JsonOutputMod[Output]{},
	}
}
